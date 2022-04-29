package assert

import (
	"fmt"
)

func Equal[T comparable](errorf errorfer, expected, actual T, msgAndArgs ...any) bool {
	if expected == actual {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] != [%#v]", expected, actual), msgAndArgs...)
}

func NotEqual[T comparable](errorf errorfer, expected, actual T, msgAndArgs ...any) bool {
	if expected != actual {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] == [%#v]", expected, actual), msgAndArgs...)
}

func Equals[T comparable](errorf errorfer, expected, actual []T, msgAndArgs ...any) bool {
	if len(expected) == len(actual) {
		equal := true

		for i, elem := range expected {
			if elem != actual[i] {
				equal = false
			}
		}

		if equal {
			return true
		}
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("[%#v] != [%#v]", expected, actual), msgAndArgs...)
}
