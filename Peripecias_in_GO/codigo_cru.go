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
	prefix            = "000000000000000000000000000000000000000000000000000000000000" // Ajustar de acordo com a carteira
	maximoCombinacoes = 65536
)

var chaves_desejadas = map[string]bool{
	"1BDyrQ6WoF8VN3g9SAS1iKZcPzFfnDVieY": true, // Carteiras alvo
	"1QCbW9HWnwQWiQqVo5exhAnmfqKRrCRsvW": true,
	"1ErZWg5cFCe4Vw5BzgfzB74VNLaXEiEkhk": true,
	"1Pie8JkxBT6MGPz9Nvi3fsPkr2D8q3GBc1": true,
}

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

// Função para gerar uma chave privada aleatória
func gerarChavePrivada() string {
	suffix := make([]byte, 2) // Mudar conforme a quantidade de bits
	_, err := rand.Read(suffix)
	if err != nil {
		log.Fatalf("Falha ao gerar chave: %v", err)
	}
	return prefix + hex.EncodeToString(suffix)
}

// Gera o hash160 da chave pública
func createPublicHash160(privKeyHex string) []byte {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	compressedPubKey := privKey.PubKey().SerializeCompressed()
	return hash160(compressedPubKey)
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
func worker(jobs <-chan struct{}) {
	defer wg.Done()
	for range jobs {
		chave := gerarChavePrivada()
		pubKeyHash := createPublicHash160(chave)
		address := encodeAddress(pubKeyHash)

		mu.Lock()
		contador++
		if chaves_desejadas[address] {
			fmt.Printf("\nEndereço encontrado: %s\nChave privada: %s\n", address, chave)
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Canal para enviar trabalhos
	jobs := make(chan struct{}, maximoCombinacoes)

	// Inicia goroutines
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go worker(jobs)
	}

	// Envia trabalhos para o canal
	for i := 0; i < maximoCombinacoes; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	// Espera todas as goroutines terminarem
	wg.Wait()

	fmt.Printf("\nTotal de chaves analisadas: %d\n", contador)
}
