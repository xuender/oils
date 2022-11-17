package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleIf() {
	fmt.Println(base.If(3 > 2, 3, 2))
	fmt.Println(base.If(3 < 2, 3, 2))

	// Output:
	// 3
	// 2
}
