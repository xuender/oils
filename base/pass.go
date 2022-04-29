package base

func Pass(error) {}

func Pass1[T any](param T, err error) T {
	return param
}

func Pass2[T1 any, T2 any](param1 T1, param2 T2, err error) (T1, T2) {
	return param1, param2
}

func Pass3[T1 any, T2 any, T3 any](param1 T1, param2 T2, param3 T3, err error) (T1, T2, T3) {
	return param1, param2, param3
}
