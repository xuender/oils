package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/times"
)

func TestNatural(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3小时", times.Natural(time.Hour*3))
	assert.Equal(t, "3分钟", times.Natural(time.Minute*3))
	assert.Equal(t, "3毫秒", times.Natural(time.Millisecond*3))
	assert.Equal(t, "3秒钟", times.Natural(time.Second*3))
	assert.Equal(t, "3小时3秒钟", times.Natural(time.Second*3+time.Hour*3))
}

func TestNaturalNano(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "3微秒", times.NaturalNano(time.Microsecond*3))
	assert.Equal(t, "2天1纳秒", times.NaturalNano(time.Hour*48+1))
}
