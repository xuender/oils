package rdb

import "errors"

var (
	ErrLock   = errors.New("is lock")
	ErrExpire = errors.New("expire time too long")
)
