package times

import (
	"strconv"
	"time"

	"github.com/xuender/oils/base"
	"golang.org/x/exp/constraints"
)

// nolint
var layouts = []string{
	"20060102",
	"2006-01-02",
	"2006-01-02 15:04:05",
	"0102",
	"2006",
	"060102",
	"06-01-02",
	"06/01/02",
	"2006/01/02",
	"060102150405",
	"060102 150405",
	"20060102150405",
	"20060102 150405",
	"06-01-02 15:04:05",
	"06/01/02 15:04:05",
	"2006/01/02 15:04:05",
}

func Parse(str string) (date time.Time, err error) {
	length := len(str)
	for _, layout := range layouts {
		if length == len(layout) {
			if date, err = time.Parse(layout, str); err == nil {
				return
			}
		}
	}

	return
}

func ParseNumber[T constraints.Integer | constraints.Float](num T) time.Time {
	const (
		len8      int64 = 100000000
		minSecond int64 = 31536000
		minMilli  int64 = 31536000000
		minMicro  int64 = 31536000000000
	)

	micro := int64(num)

	if micro < len8 {
		if newTime, err := str2time(strconv.Itoa(int(micro))); err == nil {
			return newTime
		}
	}

	if micro < minSecond {
		micro *= 1000
	}

	if micro < minMilli {
		micro *= 1000
	}

	if micro < minMicro {
		micro *= 1000
	}

	return time.UnixMicro(micro)
}

func str2time(str string) (time.Time, error) {
	switch len(str) {
	case base.There:
		return time.Parse("0102", "0"+str)
	case base.Four:
		return time.Parse("0102", str)
	case base.Five:
		return time.Parse("060102", "0"+str)
	case base.Six:
		return time.Parse("060102", str)
	case base.Eight:
		return time.Parse("20060102", str)
	}

	return time.Time{}, ErrParseError
}
