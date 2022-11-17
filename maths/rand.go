package maths

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

// RandInt 随机整数.
func RandInt(min, max int) int {
	// nolint
	return min + rand.Intn(max-min)
}

// RandBytes 随机字节码.
func RandBytes(length int) []byte {
	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		// nolint
		bytes[i] = byte(RandInt(0, 256))
	}

	return bytes
}

// RandString 随机字符串.
func RandString(length int) string {
	bytes := make([]byte, length)

	for i := 0; i < length; i++ {
		// nolint
		bytes[i] = byte(RandInt(65, 90))
	}

	return string(bytes)
}

// nolint
func init() {
	rand.Seed(time.Now().UnixNano())
}
