package syncs_test

import (
	"fmt"

	"github.com/xuender/oils/base"
	"github.com/xuender/oils/syncs"
)

func ExamplePool() {
	input := make(chan int, 10)
	output := make(chan string, 10)
	set := base.NewSet[string]()

	syncs.Pool(3, input, output, func(value, num int) string {
		return fmt.Sprintf("%d: %d*2=%d", num, value, value*2)
	})

	for i := 0; i < 100; i++ {
		input <- i

		set.Add(<-output)
	}

	close(input)
	close(output)

	fmt.Println(len(set))

	// Output:
	// 100
}
