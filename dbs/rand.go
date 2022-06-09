package dbs

import (
	"math/rand"
	"time"

	"github.com/xuender/oils/base"
)

// Rand 随机数，兼容JS数值长度.
func Rand() uint64 {
	// nolint
	return base.JSNumber(uint64(rand.Int63()))
}

// nolint
func init() {
	rand.Seed(time.Now().UnixNano())
}
