package treemap

// BytesInc 字节码加一.
func BytesInc(data []byte) []byte {
	const _max = 0xff

	if len(data) == 0 {
		return []byte{0}
	}

	res := append([]byte{}, data...)
	size := len(res)

	for i := size - 1; i >= 0; i-- {
		if res[i] != _max {
			res[i]++

			for j := i + 1; j < size; j++ {
				res[j] = 0
			}

			return res
		}
	}

	res = append(res, []byte{0xff, 0xff, 0xff}...)

	return res
}
