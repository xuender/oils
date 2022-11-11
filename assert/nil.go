package assert

import (
	"fmt"
	"reflect"
)

// IsNil 判断是否是空.
func IsNil(value any) bool {
	if value == nil {
		return true
	}

	defer func() {
		recover()
	}()

	return reflect.ValueOf(value).IsNil()
}

// Nil 判空.
func Nil(errorf errorfer, value any, msgAndArgs ...any) bool {
	if IsNil(value) {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("Expected nil, but got: %#v", value), msgAndArgs...)
}

// NotNil 非空.
func NotNil(errorf errorfer, value any, msgAndArgs ...any) bool {
	if value != nil {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, "Expected value not to be nil", msgAndArgs...)
}
