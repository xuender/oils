package cryptos_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/cryptos"
)

func ExampleEncrypt() {
	key := "key"
	str := cryptos.Encrypt("aaa", key, cryptos.AES)
	fmt.Println(cryptos.Decrypt(str, key))

	str = cryptos.Encrypt("AAA", key, cryptos.DES)
	fmt.Println(cryptos.Decrypt(str, key))

	// Output:
	// aaa <nil>
	// AAA <nil>
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	key := "key"
	src := "aaa"

	str, err := cryptos.Decrypt(cryptos.Encrypt(src, key, cryptos.AES), key)
	assert.Nil(t, err)
	assert.Equal(t, src, str)

	str, err = cryptos.Decrypt(cryptos.Encrypt(src, key, cryptos.DES), key)
	assert.Nil(t, err)
	assert.Equal(t, src, str)

	for i := 1; i <= 1000; i++ {
		src = strings.Repeat("a", i)
		str, err := cryptos.Decrypt(cryptos.Encrypt(src, key, cryptos.AES), key)
		assert.Nil(t, err)
		assert.Equal(t, src, str)
	}
}

func TestDecrypt(t *testing.T) {
	t.Parallel()

	str, err := cryptos.Decrypt("AES(A/43wTj2AVQboZZ0lNMqbw==)", "key")

	assert.Nil(t, err)
	assert.Equal(t, "aaa", str)

	_, err = cryptos.Decrypt("A/43wTj2AVQboZZ0lNMqbw==", "key")

	assert.NotNil(t, err)
}
