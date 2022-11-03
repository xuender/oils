package cryptos

import (
	"crypto/rand"
)

// Encrypt 加密.
//  txt := cryptos.Encrypt("123", "password", cryptos.AES)
//  fmt.Println(txt)
//  // AES(NR501ySw4TYFc0dI8zERfA==)
func Encrypt(str, key string, cipher Cipher) string {
	var ret []byte

	block := cipher.Block(key)
	srcBytes := Padding(str, block.BlockSize())
	tmp := make([]byte, block.BlockSize())

	for index := 0; index < len(srcBytes); index += block.BlockSize() {
		block.Encrypt(tmp, srcBytes[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	return cipher.Stringify(ret)
}

// Decrypt 解密.
//  str, _ := cryptos.Decrypt("AES(NR501ySw4TYFc0dI8zERfA==)", "password")
//  fmt.Println(str)
//  // 123
func Decrypt(src, key string) (string, error) {
	var ret []byte

	srcBytes, cipher, err := Parse(src)
	if err != nil {
		return "", err
	}

	block := cipher.Block(key)
	tmp := make([]byte, block.BlockSize())

	for index := 0; index < len(srcBytes); index += block.BlockSize() {
		block.Decrypt(tmp, srcBytes[index:index+block.BlockSize()])
		ret = append(ret, tmp...)
	}

	return UnPadding(ret)
}

func Padding(str string, blockSize int) []byte {
	cipherText := []byte(str)
	padding := getPaddingSize(cipherText, blockSize)
	padData := make([]byte, padding-1)

	_, _ = rand.Read(padData)
	// nolint
	padData = append(padData, byte(padding))

	return append(cipherText, padData...)
}

func UnPadding(cipherText []byte) (string, error) {
	ps := int(cipherText[len(cipherText)-1])

	return string(cipherText[:len(cipherText)-ps]), nil
}

func getPaddingSize(cipherText []byte, blockSize int) int {
	remainder := len(cipherText) % blockSize

	if remainder == 0 {
		return blockSize
	}

	return blockSize - remainder
}
