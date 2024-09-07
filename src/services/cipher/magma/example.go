package magma

import (
	"bufio"
	"fmt"
	"os"
)

func Example() {
	key := []byte("key1key1key1key1key1key1key1key1") // 32-bit key
	cipher := MNewCipher(key)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to encrypt: ")
	input, _ := reader.ReadString('\n')

	plaintext := []byte(input[:len(input)-1])

	encrypted := encryptText(cipher, plaintext)
	fmt.Printf("Encrypted: %x\n", encrypted)

	decrypted := decryptText(cipher, encrypted)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
