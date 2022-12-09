package syncs

// Pool Goroutine 池.
func Pool[I, O any](size int, inputs <-chan I, outputs chan<- O, call func(I, int) O) {
	for i := 0; i < size; i++ {
		go func(num int) {
			for input := range inputs {
				outputs <- call(input, num)
			}
		}(i)
	}
}
