package base

import (
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// ParseInteger 字符串转换成整数.
// nolint
func ParseInteger[T constraints.Integer](str string) (T, error) {
	if strings.ContainsRune(str, '.') {
		float, err := ParseFloat[float64](str)
		if err != nil {
			return 0, err
		}

		return Round[T](float), nil
	}

	u64, err := strconv.ParseUint(str, Ten, SixtyFour)
	// nolint
	return T(u64), err
}

// ParseFloat 字符串转换成浮点数.
func ParseFloat[T constraints.Float](str string) (T, error) {
	f64, err := strconv.ParseFloat(str, SixtyFour)
	// nolint
	return T(f64), err
}

// Itoa 数值装换成字符串.
func Itoa[T constraints.Integer | constraints.Float](num T) string {
	return strconv.Itoa(int(num))
}

// FormatFloat 浮点数格式化成字符串.
func FormatFloat[T constraints.Float | constraints.Integer](num T, prec int) string {
	return strconv.FormatFloat(float64(num), 'g', prec, SixtyFour)
}

// Round 四舍五入.
func Round[I constraints.Integer, F constraints.Float](float F) I {
	half := 0.5
	// nolint
	return I(math.Floor(float64(float) + half))
}
