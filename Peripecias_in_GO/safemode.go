package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"meugo/encoding"
	"runtime"
	"sync"
	"time"
)

const (
	prefix        = "0000000000000000000000000000000000000000000000" // Prefixo da chave
	memBufferSize = 2 * 1024 * 1024 * 1024                           // usando 2gb pro buffer
)

var chaves_desejadas = map[string]bool{
	"19vkiEajfhuZ8bs8Zu2jgmC6oqZbWqhxhG ": true, // c 69
	"1PWo3JeB9jrGwfHDNpdGK54CRas7fsVzXU ": true, // c 71
	"1JTK7s9YVYywfm5XUH7RNhHJH1LshCaRFR ": true, // c 72
}

// 4 byte xx xx xx xx

var (
	contador          int
	encontrado        bool
	mu                sync.Mutex
	wg                sync.WaitGroup
	ultimaChaveGerada string
	memBuffer         = make([]byte, memBufferSize)
)

func gerarChavePrivada() string {
	suffix := make([]byte, 9) // Tamanho do sufixo
	_, err := rand.Read(suffix)
	if err != nil {
		log.Fatalf("Falha ao gerar chave: %v", err)
	}
	chaveGerada := prefix + hex.EncodeToString(suffix)

	// Manipulação eficiente do buffer (exemplo genérico)
	copy(memBuffer[:len(suffix)], suffix)

	return chaveGerada
}

func worker(id int) {
	defer wg.Done()

	for {
		mu.Lock()
		if encontrado {
			mu.Unlock()
			return
		}
		mu.Unlock()

		chave := gerarChavePrivada()
		pubKeyHash := encoding.CreatePublicHash160(chave)
		address := encoding.EncodeAddress(pubKeyHash)

		mu.Lock()
		contador++
		ultimaChaveGerada = chave
		//fmt.Print("\n", ultimaChaveGerada) //mostrar Private keys
		//fmt.Print("\n", address) // mostrar Carteira
		if chaves_desejadas[address] {
			fmt.Printf("\n\n|--------------%s----------------|\n", address)
			fmt.Printf("|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|")
			fmt.Printf("\n|%s|\n", chave)
			encontrado = true
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}

func main() {

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(11) //numCPU//numCPU//numCPU

	// Inicia goroutines
	for i := 0; i < 11; i++ { //numCPU //numCPU //numCPU
		wg.Add(1)
		go worker(i)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			mu.Lock()
			if encontrado {
				mu.Unlock()
				break
			}
			fmt.Printf("\r N° Cpu's Disponiveis: %d | Chaves Geradas: %d | ", numCPU, contador)
			mu.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			mu.Lock()
			if encontrado {
				mu.Unlock()
				break
			}
			fmt.Printf("Ultima Chave Gerada: %s ", ultimaChaveGerada)
			fmt.Printf("  By:Ch!iNa")
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Print("\n\n|--------------------------------------------------by-Luan-BSC---|")
	fmt.Print("\n|-----------------------China-LOOP-MENU------------------------- |")
	fmt.Printf("\n|		Threads usados: %d		                 |", numCPU)
	fmt.Print("\n|		Chaves Analisadas:	", contador)
	fmt.Print("\n|________________________________________________________________|")
}
