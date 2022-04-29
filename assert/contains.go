package assert

import "fmt"

func Contains[T comparable](errorf errorfer, contains []T, sub T, msgAndArgs ...any) bool {
	for _, c := range contains {
		if c == sub {
			return true
		}
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("%#v does not contain %#v", sub, contains), msgAndArgs...)
}
