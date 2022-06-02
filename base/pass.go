package base

// Pass 忽略 error.
func Pass(error) {}

// Pass1 忽略 error, 返回1个结果.
// 推荐 value, _ := FuncTwoReturn()
// 可选 FuncOneParam(base.Pass1(FuncTwoReturn())).
func Pass1[T any](param T, err error) T {
	return param
}

// Pass2 忽略 error, 返回2个结果.
func Pass2[T1, T2 any](param1 T1, param2 T2, err error) (T1, T2) {
	return param1, param2
}

// Pass3 忽略 error, 返回3个结果.
func Pass3[T1, T2, T3 any](param1 T1, param2 T2, param3 T3, err error) (T1, T2, T3) {
	return param1, param2, param3
}
