package times_test

import (
	"fmt"
	"time"

	"github.com/xuender/oils/times"
)

func ExampleDebounced() {
	count := 0
	debouncedFn := times.Debounced(func() {
		count++
	}, 10*time.Millisecond)

	for i := 0; i < 100; i++ {
		debouncedFn()
	}

	time.Sleep(20 * time.Millisecond)

	for i := 0; i < 100; i++ {
		debouncedFn()
	}

	time.Sleep(20 * time.Millisecond)

	fmt.Println(count)

	// Output:
	// 2
}
