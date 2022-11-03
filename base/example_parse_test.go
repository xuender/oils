package base_test

import (
	"fmt"

	"github.com/xuender/oils/base"
)

func ExampleParseInteger() {
	// 字符串转换成整数
	fmt.Println(base.Must1(base.ParseInteger[int]("3")))

	// Output:
	// 3
}

func ExampleParseFloat() {
	fmt.Println(base.Must1(base.ParseFloat[float32]("3.14")))

	// Output:
	// 3.14
}

func ExampleItoa() {
	fmt.Println(base.Itoa(3))
	fmt.Println(base.Itoa(3.14))

	// Output:
	// 3
	// 3
}

func ExampleFormatFloat() {
	fmt.Println(base.FormatFloat(3, 3))
	fmt.Println(base.FormatFloat(3.14, 3))

	// Output:
	// 3
	// 3.14
}
