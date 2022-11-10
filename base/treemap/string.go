package treemap

// String 键值 string.
type String[V any] struct {
	TreeMap[string, V]
}

func NewString[V any](notFoundValue V) *String[V] {
	return &String[V]{
		TreeMap[string, V]{notFoundKey: "", notFoundValue: notFoundValue},
	}
}

// Prefix 前缀.
func (p *String[V]) Prefix(iterator ItemIterator[string, V], prefix string) {
	p.Range(iterator, prefix, PrefixNext(prefix))
}

// PrefixDesc 倒叙前缀.
func (p *String[V]) PrefixDesc(iterator ItemIterator[string, V], prefix string) {
	p.RangeDesc(iterator, prefix, PrefixNext(prefix))
}

// PrefixNext 字符串前缀下一个.
func PrefixNext(prefix string) string {
	if prefix == "" {
		return string([]byte{0})
	}

	bytes := []byte(prefix)
	bytes[len(bytes)-1]++

	return string(bytes)
}
