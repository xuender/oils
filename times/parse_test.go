package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/times"
)

func TestParse(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	tests := [][]string{
		{"19780321", "20060102"},
		{"780321", "060102"},
		{"1978-03-21", "2006-01-02"},
		{"1978-03-21 10:21:34", "2006-01-02 15:04:05"},
	}

	for _, tt := range tests {
		day := base.Must1(times.Parse(tt[0]))
		want := base.Must1(time.Parse(tt[1], tt[0]))
		ass.Equal(day, want)
	}

	_, err := times.Parse("11")
	ass.NotNil(err)
}

func TestParse_Error(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	_, err := times.Parse("23424fasdkfjk")
	is.NotNil(err)
}

func TestParseNumber(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	ass.Equal(2022, times.ParseNumber(1648777840.2672896).Year())
	ass.Equal(2022, times.ParseNumber(1648777840000).Year())
	ass.Equal(2022, times.ParseNumber(1648777840).Year())
	ass.Equal(2022, times.ParseNumber(1648777).Year())
	ass.Equal(2022, times.ParseNumber(20220401).Year())
	ass.Equal(2022, times.ParseNumber(220401).Year())
	ass.Equal(2002, times.ParseNumber(20401).Year())
	ass.Equal(time.Month(11), times.ParseNumber(1101).Month())
	ass.Equal(time.Month(4), times.ParseNumber(401).Month())
}
