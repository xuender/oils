package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleThere() {
	fmt.Println(base.One)
	fmt.Println(base.Two)
	fmt.Println(base.Three)
	fmt.Println(base.Four)
	fmt.Println(base.Five)
	fmt.Println(base.SixtyFour)
	fmt.Println(base.OneHundredTwentyEight)
	fmt.Println(base.Kilo)

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 64
	// 128
	// 1024
}
