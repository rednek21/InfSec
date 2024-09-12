package grasshopper

import (
	"bytes"
	"fmt"
)

func pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

func unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data is empty")
	}
	padLen := int(data[len(data)-1])
	if padLen > blockSize || padLen == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	return data[:len(data)-padLen], nil
}

func EncryptData(cipher *Cipher, data []byte) ([]byte, error) {
	blockSize := cipher.BlockSize()
	paddedData := pad(data, blockSize)
	encrypted := make([]byte, len(paddedData))
	for i := 0; i < len(paddedData); i += blockSize {
		cipher.Encrypt(encrypted[i:i+blockSize], paddedData[i:i+blockSize])
	}
	return encrypted, nil
}

func xor(dst, src1, src2 *[BlockSize]byte) {
	for i := 0; i < BlockSize; i++ {
		dst[i] = src1[i] ^ src2[i]
	}
}

func DecryptData(cipher *Cipher, data []byte) ([]byte, error) {
	blockSize := cipher.BlockSize()
	decrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += blockSize {
		cipher.Decrypt(decrypted[i:i+blockSize], data[i:i+blockSize])
	}
	return unpad(decrypted, blockSize)
}
