package maths_test

import (
	"fmt"

	"github.com/xuender/oils/maths"
)

func ExampleRound() {
	fmt.Println(maths.Round(3.005, 2))
	fmt.Println(maths.Round(3.004, 2))
	fmt.Println(maths.Round(3.0006, 3))

	// Output:
	// 3.01
	// 3
	// 3.001
}
