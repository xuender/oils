package base

// Conversion 类型装换是否成功.
func Conversion[T any](obj any, must bool) T {
	ret, success := obj.(T)

	if must && !success {
		panic(ErrConversion)
	}

	return ret
}
