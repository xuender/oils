package cryptos

import "errors"

var (
	ErrUnpaddingLength = errors.New("invalid unpadding length")
	ErrNoEncrypt       = errors.New("no encrypt")
)
