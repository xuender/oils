package syncs_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/oils/syncs"
)

func ExamplePool() {
	pool := syncs.NewPool(10, func(value, num int) string {
		time.Sleep(time.Millisecond)

		return fmt.Sprintf("%d: %d*2=%d", num, value, value*2)
	})

	outputs := pool.Post(lo.Range(100))

	fmt.Println(len(outputs))

	// Output:
	// 100
}

func ExamplePool_context() {
	pool := syncs.NewPool(10, func(input lo.Tuple2[context.Context, int], num int) int {
		time.Sleep(time.Millisecond)

		return input.B * input.B
	})

	inputs := lo.Map(lo.Range(100), func(num, _ int) lo.Tuple2[context.Context, int] {
		return lo.T2(context.Background(), num)
	})
	outputs := pool.Post(inputs)

	fmt.Println(len(outputs))

	// Output:
	// 100
}

func ExamplePool_error() {
	pool := syncs.NewPool(10, func(value, num int) lo.Tuple2[int, error] {
		time.Sleep(time.Millisecond)

		if value == 0 {
			// nolint
			return lo.T2(0, errors.New("divide by zero"))
		}

		return lo.T2[int, error](100/value, nil)
	})

	outputs := pool.Post(lo.Range(100))

	fmt.Println(len(outputs))

	// Output:
	// 100
}
