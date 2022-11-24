package base

import (
	"fmt"
	"strings"
)

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
}

func NewMapSameValue[K comparable, V any](value V, keys ...K) Map[K, V] {
	ret := Map[K, V]{}

	for _, key := range keys {
		ret[key] = value
	}

	return ret
}

func (p Map[K, V]) Has(key K) bool {
	_, has := p[key]

	return has
}

func (p Map[K, V]) Any(keys ...K) bool {
	for _, k := range keys {
		if _, has := p[k]; has {
			return true
		}
	}

	return false
}

func (p Map[K, V]) All(keys ...K) bool {
	for _, k := range keys {
		if _, has := p[k]; !has {
			return false
		}
	}

	return true
}

func (p Map[K, V]) Del(keys ...K) {
	for _, k := range keys {
		delete(p, k)
	}
}

func (p Map[K, V]) DelMap(elems ...Map[K, V]) {
	for _, elem := range elems {
		for k := range elem {
			delete(p, k)
		}
	}
}

// ValuesByKey 根据key获取value.
func (p Map[K, V]) ValuesByKeys(keys []K) []V { return Values(p, keys) }

func (p Map[K, V]) Put(elems ...Map[K, V]) {
	for _, elem := range elems {
		for k, v := range elem {
			p[k] = v
		}
	}
}

// String 字符串.
func (p Map[K, V]) String() string {
	var builder strings.Builder

	builder.WriteString("Map[")

	has := false
	for key, value := range p {
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
