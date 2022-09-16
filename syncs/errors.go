package syncs

import "errors"

var (
	ErrSizeLessZero = errors.New("size less than 0")
	ErrLimit        = errors.New("limit time")
)
