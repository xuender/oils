package assert

func True(errorf errorfer, value bool, msgAndArgs ...any) bool {
	if value {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, "Should be true", msgAndArgs...)
}

func False(errorf errorfer, value bool, msgAndArgs ...interface{}) bool {
	if !value {
		return true
	}

	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	return Fail(errorf, "Should be false", msgAndArgs...)
}
