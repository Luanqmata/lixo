tipo 1 :
prefix        = "00000000000000000000000000000000000000000000000"
func gerarChavePrivada() string {
	suffix := make([]byte, 9) // se for carteira impar coloca 1 byte a mais e no encode ajusta com o tamanho de caracteres da key
	_, err := rand.Read(suffix)
	if err != nil {
		log.Fatalf("Falha ao gerar chave: %v", err)
	}
	chaveGerada := prefix + hex.EncodeToString(suffix)[:17] //linha usada para gerar somente 17 caracteres

	// Manipulação eficiente do buffer (exemplo genérico)
	copy(memBuffer[:len(suffix)], suffix)

	return chaveGerada
}

tipo 2:
func gerarChavePrivada() string {
	rangeStr, existe := ranges[tamanhoChave]
	if !existe {
		log.Fatalf("Tamanho de chave %d não suportado.", tamanhoChave)
	}

	valores := strings.Split(rangeStr, "-")
	minRange := new(big.Int)
	minRange.SetString(valores[0], 16)
	maxRange := new(big.Int)
	maxRange.SetString(valores[1], 16)

	var chaveGerada *big.Int
	for {
		// Gerar um número aleatório dentro do intervalo especificado
		chaveGerada, _ = rand.Int(rand.Reader, new(big.Int).Sub(maxRange, minRange))
		chaveGerada.Add(chaveGerada, minRange)

		// Verifica se a chave gerada está dentro do intervalo permitido
		if chaveGerada.Cmp(minRange) >= 0 && chaveGerada.Cmp(maxRange) <= 0 {
			break
		}
	}

	// Converter a chave gerada para string hexadecimal
	chaveHex := hex.EncodeToString(chaveGerada.Bytes())

	// Adicionar zeros à esquerda para completar 64 caracteres
	if len(chaveHex) < 64 {
		chaveHex = strings.Repeat("0", 64-len(chaveHex)) + chaveHex
	}

	return chaveHex
}
