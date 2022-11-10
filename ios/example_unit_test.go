package ios_test

import (
	"fmt"

	"github.com/xuender/oils/ios"
)

func ExampleUnit() {
	fmt.Println(ios.Unit(1, 2))
	fmt.Println(ios.Unit(100, 2))
	fmt.Println(ios.Unit(1024, 2))
	fmt.Println(ios.Unit(1700, 2))
	fmt.Println(ios.Unit(3*ios.Meg, 2))

	// Output:
	// 1B
	// 100B
	// 1KB
	// 1.7KB
	// 3MB
}
