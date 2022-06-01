package base

func Split(str string, seps ...rune) []string {
	if len(seps) == 0 {
		return []string{str}
	}

	ret := []string{}
	slice := NewSlice[rune]()

	for _, elem := range str {
		if Has(seps, elem) {
			ret = append(ret, string(slice))
			slice.Clean()

			continue
		}

		slice.Add(elem)
	}

	ret = append(ret, string(slice))

	return ret
}
