package rsa

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Example() {
	if err := generateAndSaveKeys(); err != nil {
		log.Fatalf("Ошибка при генерации ключей: %v", err)
	}

	privateKey, err := loadPrivateKey()
	if err != nil {
		log.Fatalf("Ошибка при загрузке приватного ключа: %v", err)
	}

	publicKey, err := loadPublicKey()
	if err != nil {
		log.Fatalf("Ошибка при загрузке публичного ключа: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to encrypt: ")
	message, _ := reader.ReadString('\n')

	plaintext := []byte(message[:len(message)-1])

	ciphertext, err := encryptMessage(publicKey, plaintext)
	if err != nil {
		log.Fatalf("Ошибка при шифровании сообщения: %v", err)
	}
	fmt.Println("Сообщение успешно зашифровано")

	plaintext, err = decryptMessage(privateKey, ciphertext)
	if err != nil {
		log.Fatalf("Ошибка при расшифровании сообщения: %v", err)
	}
	fmt.Printf("Расшифрованное сообщение: %s\n", plaintext)
}
