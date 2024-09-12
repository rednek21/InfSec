package grasshopper

import (
	"bufio"
	"fmt"
	"os"
)

func Example() {
	key := []byte("key1key1key1key1key1key1key1key1") // 32-bit key
	cipher := NewCipher(key)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text to encrypt: ")
	plaintext, _ := reader.ReadString('\n')
	plaintext = plaintext[:len(plaintext)-1]

	encrypted, err := EncryptData(cipher, []byte(plaintext))
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

	decrypted, err := DecryptData(cipher, encrypted)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
