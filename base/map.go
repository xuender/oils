package base

import (
	"fmt"
	"strings"
)

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (m Map[K, V]) Has(key K) bool {
	_, has := m[key]

	return has
}

func (m Map[K, V]) Any(keys ...K) bool {
	for _, k := range keys {
		if _, has := m[k]; has {
			return true
		}
	}

	return false
}

func (m Map[K, V]) All(keys ...K) bool {
	for _, k := range keys {
		if _, has := m[k]; !has {
			return false
		}
	}

	return true
}

func (m Map[K, V]) Del(keys ...K) {
	for _, k := range keys {
		delete(m, k)
	}
}

func (m Map[K, V]) DelMap(elems ...Map[K, V]) {
	for _, elem := range elems {
		for k := range elem {
			delete(m, k)
		}
	}
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func (m Map[K, V]) Put(elems ...Map[K, V]) {
	for _, elem := range elems {
		for k, v := range elem {
			m[k] = v
		}
	}
}

// String 字符串.
func (m Map[K, V]) String() string {
	var builder strings.Builder

	builder.WriteString("Map[")

	has := false
	for key, value := range m {
		if has {
			builder.WriteString(fmt.Sprintf(" %v:%v", key, value))

			continue
		}

		builder.WriteString(fmt.Sprintf("%v:%v", key, value))

		has = true
	}

	builder.WriteString("]")

	return builder.String()
}
