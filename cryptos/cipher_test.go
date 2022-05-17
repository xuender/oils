package cryptos_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/cryptos"
)

func TestParse(t *testing.T) {
	t.Parallel()

	src, cip, err := cryptos.Parse("AES(MTIz)")

	assert.Nil(t, err)
	assert.Equals(t, []byte("123"), src)
	assert.Equal(t, cryptos.AES, cip)

	_, _, err = cryptos.Parse("AES(123")
	assert.NotNil(t, err)

	_, _, err = cryptos.Parse("AES123)")
	assert.NotNil(t, err)

	_, _, err = cryptos.Parse("ABC(123)")
	assert.NotNil(t, err)

	_, _, err = cryptos.Parse("ABC)123(")
	assert.NotNil(t, err)
}

func TestCipher_Stringify(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "AES(MTIz)", cryptos.AES.Stringify([]byte("123")))
}

func TestIsEncrypt(t *testing.T) {
	t.Parallel()

	assert.True(t, cryptos.IsEncrypt("AES(MTIz)"))
}
