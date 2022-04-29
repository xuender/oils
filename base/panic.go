package base

// Panic 恐慌.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// Panic1 恐慌1.
func Panic1[T any](t T, err error) T {
	Panic(err)

	return t
}

// Panic2 恐慌2.
func Panic2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	Panic(err)

	return t1, t2
}

// Panic3 恐慌3.
func Panic3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	Panic(err)

	return t1, t2, t3
}
