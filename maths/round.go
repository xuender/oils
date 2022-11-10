package maths

import (
	"math"
)

// Round 四舍五入.
func Round(num float64, prec int) float64 {
	pow := math.Pow10(prec)

	return math.Trunc((num+0.5/pow)*pow) / pow
}
