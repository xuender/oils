package base

import (
	"fmt"
)

// Recover 异常捕捉.
func Recover(call func(error)) {
	if rec := recover(); rec != nil {
		switch err := rec.(type) {
		case error:
			call(err)
		default:
			// nolint
			call(fmt.Errorf("%v", err))
		}
	}
}
