package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"meugo/crypto/base58"
	"runtime"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

const (
	prefix            = "000000000000000000000000000000000000000000000000000000000000" // mudar de acordo com a carteira
	maximoCombinacoes = 65536
)

var chaves_desejadas = map[string]bool{
	"1BDyrQ6WoF8VN3g9SAS1iKZcPzFfnDVieY": true, // mudar todas as carteiras conforme o que procura
	"1QCbW9HWnwQWiQqVo5exhAnmfqKRrCRsvW": true,
	"1ErZWg5cFCe4Vw5BzgfzB74VNLaXEiEkhk": true,
	"1Pie8JkxBT6MGPz9Nvi3fsPkr2D8q3GBc1": true,
}

var contador int
var wg sync.WaitGroup
var mu sync.Mutex

// Função para gerar uma chave privada aleatória
func gerarChavePrivada() string {
	suffix := make([]byte, 2) // mudar de acordo com a quantidade de bits
	_, err := rand.Read(suffix)
	if err != nil {
		log.Fatalf("Falha ao gerar chave: %v", err)
	}
	chaveGerada := prefix + hex.EncodeToString(suffix)
	return chaveGerada
}

// Função que gera WIF
func generateWif(privKeyHex string) string {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	extendedKey := append([]byte{byte(0x80)}, privKeyBytes...)
	extendedKey = append(extendedKey, byte(0x01))

	firstSHA := sha256.Sum256(extendedKey)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	finalKey := append(extendedKey, checksum...)
	wif := base58.Encode(finalKey)
	return wif
}

// Gera o hash160 da chave pública
func createPublicHash160(privKeyHex string) []byte {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	compressedPubKey := privKey.PubKey().SerializeCompressed()

	pubKeyHash := hash160(compressedPubKey)
	return pubKeyHash
}

// Hash SHA256 seguido de RIPEMD160
func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}

// Codifica o endereço a partir do hash da chave pública
func encodeAddress(pubKeyHash []byte) string {
	version := byte(0x00)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := doubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}

// Faz o duplo hash SHA256
func doubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

// Função que será executada em paralelo por cada goroutine
func worker(id int, jobs <-chan struct{}) {
	defer wg.Done()

	for range jobs {
		chave := gerarChavePrivada()
		pubKeyHash := createPublicHash160(chave)
		address := encodeAddress(pubKeyHash)

		mu.Lock()
		contador++
		if chaves_desejadas[address] {
			fmt.Printf("\n\n|--------------%s----------------|\n", address)
			fmt.Printf("|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|")
			fmt.Printf("\n|%s|\n", chave)
			mu.Unlock()
			return // Encerra o programa quando a chave correta é encontrada
		}
		mu.Unlock()
	}
}

func main() {
	// Determina o número de CPUs disponíveis e usa todos
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Canal para enviar trabalhos (geração de chaves)
	jobs := make(chan struct{}, maximoCombinacoes)

	// Inicia goroutines
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go worker(i, jobs)
	}

	// Envia trabalhos para o canal
	for i := 0; i < maximoCombinacoes; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	// Espera todas as goroutines terminarem
	wg.Wait()

	fmt.Println("\n                              ______")
	fmt.Println("                             |      |")
	fmt.Println("                             |OOPS! |")
	fmt.Println("                             |WALLET|")
	fmt.Println("                             |FOUND!|")
	fmt.Println("                             |______|")
	fmt.Print("|--------------------------------------------------by-Luan-BSC---|")
	fmt.Print("\n|-----------------------China-LOOP-MENU------------------------- |")
	fmt.Printf("\n|		Threads usados: %d		                 |", numCPU)
	fmt.Print("\n|		Maximo de possibilidades: ", maximoCombinacoes)
	fmt.Print("\n|		Chaves Analisadas:	", contador)
	fmt.Print("\n|________________________________________________________________|")
}
