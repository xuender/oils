package cryptos_test

import (
	"fmt"

	"github.com/xuender/oils/cryptos"
)

func Example() {
	password := "password"
	txt := cryptos.Encrypt("123", password, cryptos.DES)

	// fmt.Println(txt)
	fmt.Println(cryptos.IsEncrypt(txt))
	fmt.Println(cryptos.Decrypt(txt, password))

	// Output:
	// true
	// 123 <nil>
}
