package assert

type tHelper interface {
	Helper()
}

type errorfer interface {
	Errorf(format string, args ...interface{})
}
