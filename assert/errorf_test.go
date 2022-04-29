package assert_test

type errorfer struct{}

func (p *errorfer) Errorf(format string, args ...interface{}) {}
func (p *errorfer) Helper()                                   {}
