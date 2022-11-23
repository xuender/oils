package syncs_test

import (
	"testing"

	"github.com/xuender/oils/base"
)

func BenchmarkIfElse(b *testing.B) {
	call := func(index int) {
		var num int

		hund := base.Hundred

		if index < base.Ten {
			num += index
		} else {
			num += hund
		}

		base.Pass1(num, nil)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		call(n)
	}
}

func BenchmarkIf(b *testing.B) {
	call := func(index int) {
		var num int

		hund := base.Hundred

		if index < base.Ten {
			hund = index
		}

		num += hund

		base.Pass1(num, nil)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		call(n)
	}
}
