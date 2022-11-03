package cryptos_test

import (
	"fmt"

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
