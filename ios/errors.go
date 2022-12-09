package ios

import "errors"

var (
	ErrTimeout   = errors.New("time out")
	ErrChanClose = errors.New("channel close")
)
