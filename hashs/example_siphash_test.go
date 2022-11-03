package hashs_test

import (
	"fmt"

	"github.com/xuender/oils/hashs"
)

func ExampleSipHash128() {
	fmt.Println(hashs.SipHash128([]byte("123")))

	// Output:
	// 8693645449139915215 11618447955228391416
}

func ExampleSipHash64() {
	fmt.Println(hashs.SipHash64([]byte("123")))

	// Output:
	// 9379172312344772015
}

func ExampleSipHashNumber() {
	fmt.Println(hashs.SipHashNumber([]byte("123")))
	// output:
	// 2677888159399343
}

func ExampleSipHashHex() {
	fmt.Println(hashs.SipHashHex([]byte("123")))
	// output:
	// 822983866c7d3daf
}
