package base

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{}
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

// Clear 清除.
func (p Map[K, V]) Clear() { maps.Clear(p) }

// Clone 克隆.
func (p Map[K, V]) Clone() Map[K, V] { return maps.Clone(p) }

// Copy 复制.
func (p Map[K, V]) Copy(dst Map[K, V]) { maps.Copy(dst, p) }

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

// Keys 键切片.
func (p Map[K, V]) Keys() []K { return maps.Keys(p) }

// Values 值切片.
func (p Map[K, V]) Values() []V { return maps.Values(p) }

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
