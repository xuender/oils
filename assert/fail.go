package assert

import (
	"fmt"
	"strings"
)

func Fail(errorf errorfer, failureMessage string, msgAndArgs ...any) bool {
	if h, ok := errorf.(tHelper); ok {
		h.Helper()
	}

	if len(msgAndArgs) > 0 {
		// nolint: asasalint
		errorf.Errorf("%s: %s\n", Message(msgAndArgs), failureMessage)
	} else {
		errorf.Errorf("%s\n", failureMessage)
	}

	return false
}

func Message(args ...any) string {
	two := 2

	switch len(args) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", args[0])
	case two:
		return fmt.Sprintf("%v(%#v)", args...)
	default:
		return fmt.Sprintf("%v(%#v"+strings.Repeat(", %#v", len(args)-two)+")", args...)
	}
}
