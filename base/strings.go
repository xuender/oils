package base

import (
	"strings"
	"unicode"
)

func SplitSpace(str string) []string {
	data := []string{}

	for {
		str = strings.TrimSpace(str)
		index := strings.IndexFunc(str, unicode.IsSpace)

		if index < 0 {
			break
		}

		data = append(data, str[:index])
		str = str[index:]
	}

	if len(str) > 0 {
		data = append(data, str)
	}

	return data
}
