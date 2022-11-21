package oss

import (
	"fmt"
	"sync"
)

// nolint: gochecknoglobals
var _unique = NewUnique()

// Unique 唯一.
type Unique struct {
	id   int
	lock sync.Mutex
}

// NewUnique 新建唯一.
func NewUnique() *Unique {
	return &Unique{}
}

// ID 唯一整数.
func (p *Unique) ID() int {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.id++

	return p.id
}

// StringID 唯一字符串.
func (p *Unique) StringID(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, p.ID())
}

// UniqueID 唯一ID.
func UniqueID() int {
	return _unique.ID()
}

// UniqueString 唯一字符串.
func UniqueString(prefix string) string {
	return _unique.StringID(prefix)
}
