package base

import (
	"errors"
	"strings"
)

var ErrConversion = errors.New("conversion bool is false")

type sliceErrors []error

func Errors(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	ret := Filter(errs, func(err error) bool { return err != nil })

	switch len(ret) {
	case 0:
		return nil
	case 1:
		return ret[0]
	default:
		return sliceErrors(ret)
	}
}

func (p sliceErrors) Error() string {
	builder := &strings.Builder{}

	for index, err := range p {
		if index > 0 {
			builder.WriteString(", ")
		}

		builder.WriteString(err.Error())
	}

	return builder.String()
}
