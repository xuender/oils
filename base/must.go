package base

// Must 必须.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// Panic 恐慌.
//
// As of v0.1.9, this function simply calls Must.
func Panic(err error) {
	Must(err)
}

// Must1 必须1.
func Must1[T any](t T, err error) T {
	Must(err)

	return t
}

// Panic1 恐慌1.
//
// As of v0.1.9, this function simply calls Must1.
func Panic1[T any](t T, err error) T {
	return Must1(t, err)
}

// Must2 必须2.
func Must2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	Must(err)

	return t1, t2
}

// Must3 必须3.
func Must3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	Must(err)

	return t1, t2, t3
}
