package hashs

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/dchest/siphash"
	"github.com/xuender/oils/base"
)

const (
	key0 = 506097522914230528
	key1 = 1084818905618843912
)

// SipHash32 32位哈希.
func SipHash32(data []byte) uint32 {
	sum := make([]byte, base.Eight)
	binary.BigEndian.PutUint64(sum, SipHash64(data))

	return binary.BigEndian.Uint32(sum[0:base.Four])
}

// SipHash128 哈希.
func SipHash128(data []byte) (uint64, uint64) {
	return siphash.Hash128(key0, key1, data)
}

// SipHash64 哈希，和Google Guava的sipHash24相同.
func SipHash64(data []byte) uint64 {
	return siphash.Hash(key0, key1, data)
}

// SipHashNumber 兼容JS, 53位长度.
func SipHashNumber(data []byte) uint64 {
	return base.JSNumber(siphash.Hash(key0, key1, data))
}

// SipHashHex 字符串.
func SipHashHex(data []byte) string {
	sum := siphash.Hash(key0, key1, data)
	b8 := make([]byte, base.Eight)

	binary.BigEndian.PutUint64(b8, sum)

	return hex.EncodeToString(b8)
}
