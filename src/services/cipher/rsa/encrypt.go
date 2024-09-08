package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// encryptMessage шифрует сообщение с использованием публичного ключа
func encryptMessage(publicKey *rsa.PublicKey, message []byte) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при шифровании: %w", err)
	}
	return ciphertext, nil
}
