showProgress := false // ou true, se você quiser habilitar

if showProgress {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			mu.Lock()
			if encontrado {
				mu.Unlock()
				break
			}
			duration := time.Since(startTime).Seconds()
			chavesPorSegundo := float64(contador) / duration
			fmt.Printf("\r  N° Threads Usados: %d | Chaves Geradas: %d | Chaves Por Segundo: %.2f |", numThreads, contador, chavesPorSegundo)
			fmt.Print("Ultima Chave Gerada: ", ultimaChaveGerada, " |")
			mu.Unlock()
		}
	}()
}
