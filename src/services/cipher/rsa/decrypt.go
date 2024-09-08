package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

// decryptMessage дешифрует сообщение с использованием приватного ключа.
func decryptMessage(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при расшифровании: %w", err)
	}
	return plaintext, nil
}
