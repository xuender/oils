package ordered

// Times 调用 iteratee num 次，调用返回的结果存入到切片返回.
func Times[T any](num int, iteratee func(int) T) []T {
	res := make([]T, num)

	for i := 0; i < num; i++ {
		res[i] = iteratee(i)
	}

	return res
}
