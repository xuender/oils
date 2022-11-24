package base

import (
	"fmt"
)

// Recover 异常捕捉.
func Recover(yield func(error)) {
	if rec := recover(); rec != nil {
		switch err := rec.(type) {
		case error:
			yield(err)
		default:
			// nolint
			yield(fmt.Errorf("%v", err))
		}
	}
}
