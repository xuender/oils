package dbs_test

import (
	"fmt"

	"github.com/xuender/oils/dbs"
)

func ExampleRand() {
	num := dbs.Rand()

	fmt.Println(num > 0)

	// Output:
	// true
}
