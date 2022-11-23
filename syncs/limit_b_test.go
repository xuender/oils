package syncs_test

import "testing"

func BenchmarkIfElse(b *testing.B) {
	call := func(i int) {
		var n int
		a := 100

		if i < 10000 {
			n += i
		} else {
			n += a
		}
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		call(n)
	}
}

func BenchmarkIf(b *testing.B) {
	call := func(i int) {
		var n int
		a := 100

		if i < 10000 {
			a = i
		}

		n += a
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		call(n)
	}
}
