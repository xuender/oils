package ios_test

import (
	"fmt"

	"github.com/xuender/oils/ios"
)

func ExampleWrite() {
	data := make(chan int, 2)
	fmt.Println(ios.Write(1, data))
	close(data)
	fmt.Println(ios.Write(1, data))

	// Output:
	// <nil>
	// channel close
}

func ExampleRead() {
	data := make(chan int, 1)

	data <- 1
	fmt.Println(ios.Read(data))
	close(data)
	fmt.Println(ios.Read(data))

	// Output:
	// 1 <nil>
	// 0 channel close
}
