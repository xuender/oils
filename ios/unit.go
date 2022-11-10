package ios

import (
	"fmt"

	"github.com/xuender/oils/base"
	"github.com/xuender/oils/maths"
)

const (
	_ = 1 << (10 * iota)
	Kilo
	Meg
	Giga
	Tera
	Peta
	Exa
)

func Unit(size int64, prec int) string {
	switch {
	case size < Kilo:
		return fmt.Sprintf("%dB", size)
	case size < Meg:
		return fmt.Sprintf("%sKB", base.FormatFloat(maths.Round(float64(size)/Kilo, prec), prec))
	case size < Giga:
		return fmt.Sprintf("%sMB", base.FormatFloat(maths.Round(float64(size)/Meg, prec), prec))
	case size < Tera:
		return fmt.Sprintf("%sGB", base.FormatFloat(maths.Round(float64(size)/Giga, prec), prec))
	case size < Peta:
		return fmt.Sprintf("%sTB", base.FormatFloat(maths.Round(float64(size)/Tera, prec), prec))
	case size < Exa:
		return fmt.Sprintf("%sPB", base.FormatFloat(maths.Round(float64(size)/Peta, prec), prec))
	default:
		return fmt.Sprintf("%sEB", base.FormatFloat(maths.Round(float64(size)/Exa, prec), prec))
	}
}
