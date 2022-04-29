package base

import (
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

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

func ParseFloat[T constraints.Float](str string) (T, error) {
	f64, err := strconv.ParseFloat(str, SixtyFour)
	// nolint
	return T(f64), err
}

func Itoa[T constraints.Integer | constraints.Float](num T) string {
	return strconv.Itoa(int(num))
}

func FormatFloat[T constraints.Float | constraints.Integer](num T, prec int) string {
	return strconv.FormatFloat(float64(num), 'g', prec, SixtyFour)
}

// Round 四舍五入.
func Round[I constraints.Integer, F constraints.Float](float F) I {
	half := 0.5
	// nolint
	return I(math.Floor(float64(float) + half))
}
