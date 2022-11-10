package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des" // nolint
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/xuender/oils/base"
)

type Cipher int

const (
	AES Cipher = iota
	DES
)

// nolint
var ciphers = [...]Cipher{AES, DES}

func (p Cipher) String() string {
	if p == DES {
		return "DES"
	}

	return "AES"
}

func (p Cipher) Block(key string) cipher.Block {
	keyBytes := sha256.Sum256([]byte(key))

	if p == DES {
		// nolint
		return base.Must1(des.NewCipher(keyBytes[:8]))
	}

	return base.Must1(aes.NewCipher(keyBytes[:]))
}

func (p Cipher) Stringify(data []byte) string {
	return fmt.Sprintf("%s(%s)", p.String(), base64.StdEncoding.EncodeToString(data))
}

// IsEncrypt 是否加密.
func IsEncrypt(str string) bool {
	_, _, err := Parse(str)

	return err == nil
}

// Parse 解析密文, 返回数据和加密算法.
func Parse(str string) (src []byte, cipher Cipher, err error) { // nolint
	start := strings.Index(str, "(")
	end := strings.LastIndex(str, ")")

	if start < 0 || end != len(str)-1 || start > end {
		err = ErrNoEncrypt

		return
	}

	err = ErrNoEncrypt

	for _, cip := range ciphers {
		if strings.HasPrefix(str, cip.String()) {
			cipher = cip
			err = nil
		}
	}

	if err != nil {
		return
	}

	src, err = base64.StdEncoding.DecodeString(str[start+1 : end])

	return
}
