// codigo BASE para o modo procura Varredura

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/crypto/ripemd160"
)

var (
	wg                sync.WaitGroup
	encontrado        bool
	contador          int64 // Alterado para int64 para uso seguro com atomic
	enderecoDesejado  string
	tamanhoChave      int
	ranges            map[int]string
	chaveAtual        *big.Int
	mu                sync.Mutex
	startTime         time.Time
	progressoFile     = "progresso.json"
	salvarProgressoCh = make(chan bool, 1) // Canal para controlar o salvamento do progresso
)

func init() {
	ranges = make(map[int]string)
}

func carregarRangesDoArquivo(nomeDoArquivo string) {
	file, err := os.Open(nomeDoArquivo)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de ranges: %v", err)
	}
	defer file.Close()

	for {
		var bits int
		var rangeStr string
		_, err := fmt.Fscanf(file, "%d %s\n", &bits, &rangeStr)
		if err != nil {
			break
		}
		ranges[bits] = rangeStr
	}
}

func carregarProgresso() {
	// Verifica se o arquivo de progresso existe
	if _, err := os.Stat(progressoFile); os.IsNotExist(err) {
		return
	}

	// Carrega o progresso do arquivo
	data, err := ioutil.ReadFile(progressoFile)
	if err != nil {
		log.Fatalf("Erro ao carregar o progresso: %v", err)
	}

	var progresso map[string]string
	err = json.Unmarshal(data, &progresso)
	if err != nil {
		log.Fatalf("Erro ao decodificar o progresso: %v", err)
	}

	// Se a chave atual foi salva, carregue-a
	if chaveHex, ok := progresso[fmt.Sprintf("%d", tamanhoChave)]; ok {
		chaveAtual = new(big.Int)
		// Converte a string hexadecimal de volta para big.Int
		_, success := chaveAtual.SetString(chaveHex, 16)
		if !success {
			log.Fatalf("Erro ao converter a chave atual de string para big.Int: %v", chaveHex)
		}
	}
}

func salvarProgresso() {
	mu.Lock()
	defer mu.Unlock()

	// Cria ou abre o arquivo de progresso
	file, err := os.OpenFile(progressoFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo de progresso: %v", err)
	}
	defer file.Close()

	// Serializa o progresso atual
	progresso := map[string]string{
		fmt.Sprintf("%d", tamanhoChave): chaveAtual.Text(16),
	}
	data, err := json.Marshal(progresso)
	if err != nil {
		log.Fatalf("Erro ao codificar o progresso: %v", err)
	}

	// Salva no arquivo
	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Erro ao salvar o progresso: %v", err)
	}
}

func gerarChavePrivadaSequencial(minRange, maxRange *big.Int) *big.Int {
	mu.Lock()
	defer mu.Unlock()

	// Inicializando a chave atual para o início do range
	if chaveAtual == nil {
		chaveAtual = new(big.Int).Set(minRange)
	}

	// Verifica se a chave atual está dentro do intervalo permitido
	if chaveAtual.Cmp(maxRange) > 0 {
		return nil
	}

	// Incrementa a chave atual para a próxima iteração
	chave := new(big.Int).Set(chaveAtual)
	chaveAtual.Add(chaveAtual, big.NewInt(1))

	return chave
}

func createPublicKey(chave string) string {
	privateKeyBytes, _ := hex.DecodeString(chave)
	privateKey := new(ecdsa.PrivateKey)
	privateKey.PublicKey.Curve = elliptic.P256()
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())

	// Obter a chave pública em bytes
	pubKeyBytes := elliptic.Marshal(privateKey.PublicKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)

	// Criar o hash SHA-256 da chave pública
	hash := sha256.New()
	hash.Write(pubKeyBytes)
	pubKeyHash := hash.Sum(nil)

	// Criar o hash RIPEMD-160 do hash SHA-256
	ripeHasher := ripemd160.New()
	ripeHasher.Write(pubKeyHash)
	return hex.EncodeToString(ripeHasher.Sum(nil))
}

// Função para imprimir a chave sem zeros à esquerda
func printarChaveSemZero(chave *big.Int) string {
	chaveHex := hex.EncodeToString(chave.Bytes())
	// Remove zeros à esquerda
	return strings.TrimLeft(chaveHex, "0")
}

func worker(minRange, maxRange *big.Int) {
	defer wg.Done()
	enderecoDesejado := "3ee4133d991f52fdf6a25c9834e0745ac74248a4"

	for !encontrado {
		chaveBigInt := gerarChavePrivadaSequencial(minRange, maxRange)
		if chaveBigInt == nil {
			break
		}

		chave := hex.EncodeToString(chaveBigInt.Bytes())
		pubKeyHash := createPublicKey(chave)

		atomic.AddInt64(&contador, 1) // Atualiza o contador de forma segura
		if pubKeyHash == enderecoDesejado {
			fmt.Printf("\n\n|--------------%s----------------|\n", pubKeyHash)
			fmt.Printf("|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|")
			fmt.Printf("\n|%s|\n", chave)
			encontrado = true
			break
		}

		// Salva o progresso a cada 1000 iterações
		if atomic.LoadInt64(&contador)%1000000 == 0 {
			salvarProgresso()
		}

		// Atualização da taxa de chaves por segundo
		if atomic.LoadInt64(&contador)%1000 == 0 {
			elapsedTime := time.Since(startTime).Seconds()
			keysPerSecond := float64(atomic.LoadInt64(&contador)) / elapsedTime
			// Usando printarChaveSemZero para exibir a chave atual sem zeros à esquerda
			fmt.Printf("\rChaves Geradas: %d | Velocidade: %.2f Chaves/seg | PublicHash160 : %s | Chave Atual : %s", atomic.LoadInt64(&contador), keysPerSecond, pubKeyHash, printarChaveSemZero(chaveBigInt))
		}
	}
}

func main() {
	// Capturar sinais do sistema para salvar o progresso antes de sair
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nRecebido sinal de interrupção, salvando progresso...")
		salvarProgresso()
		os.Exit(0)
	}()

	// Carregar os ranges do arquivo
	carregarRangesDoArquivo("ranges.txt")

	tamanhoChave := 67
	// Verificar se o tamanho está dentro do intervalo permitido
	rangeStr, ok := ranges[tamanhoChave]
	if !ok {
		fmt.Printf("	Tamanho de chave %d não suportado. Escolha um valor entre 1 e %d.\n", tamanhoChave, len(ranges))
		return
	}

	// Definir range
	valores := strings.Split(rangeStr, "-")
	minRange := new(big.Int)
	minRange.SetString(valores[0], 16)
	maxRange := new(big.Int)
	maxRange.SetString(valores[1], 16)

	// Carregar o progresso do arquivo JSON, se existir
	carregarProgresso()

	// Tempo de início
	startTime = time.Now()

	// Número de goroutines a serem usadas
	numWorkers := 20 // Defina o número de threads desejado

	// Iniciar goroutines para gerar chaves
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(minRange, maxRange)
	}

	// Aguardar todas as goroutines concluírem
	wg.Wait()

	// Salvar progresso final
	salvarProgresso()

	// Tempo final e saída
	elapsedTime := time.Since(startTime).Seconds()
	fmt.Printf("\nProcesso concluído.\n")
	fmt.Printf("Tempo total: %.2f segundos\n", elapsedTime)
}
