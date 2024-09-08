package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	privateKeyPath = "private_key.pem"
	publicKeyPath  = "public_key.pem"
	keySize        = 16384
)

func generateAndSaveKeys() error {
	if _, err := os.Stat(privateKeyPath); !os.IsNotExist(err) {
		fmt.Println("Ключи уже существуют.")
		return nil
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return fmt.Errorf("ошибка при генерации ключа: %w", err)
	}

	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла приватного ключа: %w", err)
	}
	defer privateKeyFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	err = pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	if err != nil {
		return fmt.Errorf("ошибка при записи приватного ключа в файл: %w", err)
	}

	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("ошибка при маршализации публичного ключа: %w", err)
	}

	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла публичного ключа: %w", err)
	}
	defer publicKeyFile.Close()

	err = pem.Encode(publicKeyFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	if err != nil {
		return fmt.Errorf("ошибка при записи публичного ключа в файл: %w", err)
	}

	fmt.Println("Ключи успешно сгенерированы и сохранены")
	return nil
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла приватного ключа: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("не удалось декодировать PEM блок")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе приватного ключа: %w", err)
	}

	return privateKey, nil
}

func loadPublicKey() (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла публичного ключа: %w", err)
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("не удалось декодировать PEM блок")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе публичного ключа: %w", err)
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("неверный тип публичного ключа")
	}

	return publicKey, nil
}
