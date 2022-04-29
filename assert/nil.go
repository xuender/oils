package assert

import "fmt"

func Nil(errorf errorfer, value any, msgAndArgs ...any) bool {
	if value == nil {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("Expected nil, but got: %#v", value), msgAndArgs...)
}

func NotNil(errorf errorfer, value any, msgAndArgs ...any) bool {
	if value != nil {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, "Expected value not to be nil", msgAndArgs...)
}
