package syncs

// Merge 多个管道内容合并.
func Merge[T any](less func(T, T) bool, chans ...chan T) <-chan T {
	out := make(chan T)
	opens := map[int]T{}

	go func() {
		for i, cha := range chans {
			if value, open := <-cha; open {
				opens[i] = value
			}
		}

		var (
			minIndex int
			min      T
		)

		for len(opens) > 0 {
			minIndex = -1

			for i, value := range opens {
				if minIndex == -1 || less(value, min) {
					minIndex = i
					min = value
				}
			}

			out <- min

			if value, open := <-chans[minIndex]; open {
				opens[minIndex] = value
			} else {
				delete(opens, minIndex)
			}
		}

		close(out)
	}()

	return out
}
