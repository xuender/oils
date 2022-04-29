package assert

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Greater[T constraints.Ordered](errorf errorfer, big, small T, msgAndArgs ...any) bool {
	if big > small {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] <= [%#v]", big, small), msgAndArgs...)
}

func GreaterOrEqual[T constraints.Ordered](errorf errorfer, big, small T, msgAndArgs ...any) bool {
	if big >= small {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] < [%#v]", big, small), msgAndArgs...)
}

func Less[T constraints.Ordered](errorf errorfer, small, big T, msgAndArgs ...any) bool {
	if small < big {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] >= [%#v]", small, big), msgAndArgs...)
}

func LessOrEqual[T constraints.Ordered](errorf errorfer, small, big T, msgAndArgs ...any) bool {
	if small <= big {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] > [%#v]", small, big), msgAndArgs...)
}
