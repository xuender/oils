package maths_test

import (
	"fmt"

	"github.com/xuender/oils/maths"
)

func ExampleRand() {
	num := maths.Rand()

	fmt.Println(num > 0)

	// Output:
	// true
}
