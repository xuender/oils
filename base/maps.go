package base

// Values 根据keys获取map的values.
func Values[M ~map[K]V, K comparable, V any](m M, keys []K) []V {
	values := make([]V, len(keys))
	for index, key := range keys {
		values[index] = m[key]
	}

	return values
}
