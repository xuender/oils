package hashs

import (
	"bytes"
	"encoding/gob"
	"reflect"

	"github.com/xuender/oils/base"
)

// Encode 转换字节码.
// Deprecated: 直接使用god.
func Encode(elem any) []byte {
	if elem == nil {
		return []byte{}
	}

	switch obj := elem.(type) {
	case string:
		return []byte(obj)
	case []byte:
		return obj
	default:
		buf := &bytes.Buffer{}
		encode := gob.NewEncoder(buf)

		base.Must(encode.Encode(elem))

		return buf.Bytes()
	}
}

// Decode 字节码解码.
// Deprecated: 直接使用god.
func Decode[T any](elem []byte) T {
	obj := new(T)

	if elem == nil {
		return *obj
	}

	switch (any)(*obj).(type) {
	case string:
		value := reflect.ValueOf(obj).Elem()
		value.SetString(string(elem))

		return *obj
	case []byte:
		value := reflect.ValueOf(obj).Elem()
		value.SetBytes(elem)

		return *obj
	default:
		buf := bytes.NewBuffer(elem)
		decoder := gob.NewDecoder(buf)
		base.Must(decoder.Decode(obj))

		return *obj
	}
}
