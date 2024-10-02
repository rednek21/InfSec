package stribog

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func Example() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку для хеширования: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Ошибка при чтении ввода: %v", err)
	}

	message = message[:len(message)-1]

	h := New(64)

	_, err = h.Write([]byte(message))
	if err != nil {
		return
	}

	hash := h.Sum(nil)

	fmt.Printf("Hash строки '%s': %s\n", message, hex.EncodeToString(hash))
}
