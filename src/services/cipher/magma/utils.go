package magma

// Шифрование текста произвольной длины
func encryptText(cipher *MCipher, plaintext []byte) []byte {
	blockSize := cipher.BlockSize()
	paddedText := padText(plaintext, blockSize) // Добавление паддинга
	encrypted := make([]byte, len(paddedText))

	for i := 0; i < len(paddedText); i += blockSize {
		cipher.Encrypt(encrypted[i:i+blockSize], paddedText[i:i+blockSize])
	}
	return encrypted
}

// Расшифрование текста произвольной длины
func decryptText(cipher *MCipher, encrypted []byte) []byte {
	blockSize := cipher.BlockSize()
	decrypted := make([]byte, len(encrypted))

	for i := 0; i < len(encrypted); i += blockSize {
		cipher.Decrypt(decrypted[i:i+blockSize], encrypted[i:i+blockSize])
	}
	return unpadText(decrypted)
}

// Добавление паддинга (для примера используем простое добавление байт)
func padText(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	paddedText := append(text, make([]byte, padding)...)
	return paddedText
}

// Удаление паддинга
func unpadText(text []byte) []byte {
	length := len(text)
	padding := int(text[length-1])
	return text[:length-padding]
}
