package base

import (
	"strings"
)

// Values 根据keys获取map的values.
func Values[M ~map[K]V, K comparable, V any](m M, keys []K) []V {
	values := make([]V, len(keys))
	for index, key := range keys {
		values[index] = m[key]
	}

	return values
}

// String2Map 字符串转换成 map.
func String2Map(split, equal string, elems ...string) map[string]string {
	ret := map[string]string{}

	for _, elem := range elems {
		for _, keyValue := range strings.Split(elem, split) {
			data := strings.Split(keyValue, equal)
			if len(data) > 1 {
				ret[data[0]] = data[1]

				continue
			}

			ret[data[0]] = ""
		}
	}

	return ret
}

// Tag2Map stuct 的 Tag 转换成 map.
func Tag2Map(tag string) map[string]string {
	return String2Map(",", "=", tag)
}
