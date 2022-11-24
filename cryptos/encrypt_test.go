package cryptos_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/cryptos"
)

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

func TestPadding(t *testing.T) {
	t.Parallel()

	data := cryptos.Padding("123", 4)
	assert.Equal(t, 4, len(data))
}

func TestUnPadding(t *testing.T) {
	t.Parallel()

	data := cryptos.Padding("123", 14)
	str, err := cryptos.UnPadding(data)

	assert.Nil(t, err)
	assert.Equal(t, "123", str)
}
