package assert

import "fmt"

type panicFunc func()

func Panics(errorf errorfer, call panicFunc, msgAndArgs ...any) bool {
	ret := panics(call)
	if ret != nil {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", call, ret), msgAndArgs...)
}

func PanicsWithError(errorf errorfer, errString string, call panicFunc, msgAndArgs ...any) bool {
	ret := panics(call)
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
		fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v not %#v", call, ret, errString),
		msgAndArgs...,
	)
}

// nolint: nonamedreturns
func panics(call panicFunc) (ret any) {
	defer func() {
		ret = recover()
	}()

	call()

	return
}
