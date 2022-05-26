package base

// Must 恐慌.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Must1 恐慌1.
func Must1[T any](t T, err error) T {
	Must(err)

	return t
}

// Must2 恐慌2.
func Must2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	Must(err)

	return t1, t2
}

// Must3 恐慌3.
func Must3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	Must(err)

	return t1, t2, t3
}
