package base

// ConversionPass 忽略类型装换是否成功.
func ConversionPass[T any](obj any) T {
	ret, _ := obj.(T)

	return ret
}

// ConversionMust 类型装换必须成功.
func ConversionMust[T any](obj any) T {
	ret, ok := obj.(T)
	if !ok {
		panic(ErrConversion)
	}

	return ret
}
