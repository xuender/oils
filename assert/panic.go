package assert

import "fmt"

type panicFunc func()

func Panics(errorf errorfer, yield panicFunc, msgAndArgs ...any) bool {
	ret := panics(yield)
	if ret != nil {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", yield, ret), msgAndArgs...)
}

func PanicsWithError(errorf errorfer, errString string, yield panicFunc, msgAndArgs ...any) bool {
	ret := panics(yield)
	switch obj := ret.(type) {
	case string:
		if obj == errString {
			return true
		}
	case error:
		if obj.Error() == errString {
			return true
		}
	default:
		if fmt.Sprintf("%#v", obj) == errString {
			return true
		}
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf,
		fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v not %#v", yield, ret, errString),
		msgAndArgs...,
	)
}

// nolint: nonamedreturns
func panics(yield panicFunc) (ret any) {
	defer func() {
		ret = recover()
	}()

	yield()

	return
}
