package times

import (
	"fmt"
	"time"
)

// Natural 自然语言显示.
func Natural(d time.Duration) string {
	return NaturalNano(d - d%time.Millisecond)
}

// NaturalNano 自然语言显示精确到纳秒.
func NaturalNano(dur time.Duration) string {
	ret := ""

	if h := int(dur.Hours()); h > 0 {
		w := 24
		if day := h / w; day > 0 {
			ret += fmt.Sprintf("%d天", day)
			dur -= time.Hour * time.Duration(day*w)
		}

		if h := int(dur.Hours()); h > 0 {
			ret += fmt.Sprintf("%d小时", h)
			dur -= time.Hour * time.Duration(h)
		}
	}

	if m := int(dur.Minutes()); m > 0 {
		ret += fmt.Sprintf("%d分钟", m)
		dur -= time.Minute * time.Duration(m)
	}

	if s := int(dur.Seconds()); s > 0 {
		ret += fmt.Sprintf("%d秒钟", s)
		dur -= time.Second * time.Duration(s)
	}

	if m := int(dur.Milliseconds()); m > 0 {
		ret += fmt.Sprintf("%d毫秒", m)
		dur -= time.Millisecond * time.Duration(m)
	}

	if m := int(dur.Microseconds()); m > 0 {
		ret += fmt.Sprintf("%d微秒", m)
		dur -= time.Microsecond * time.Duration(m)
	}

	if dur.Nanoseconds() > 0 {
		ret += fmt.Sprintf("%d纳秒", dur.Nanoseconds())
	}

	return ret
}
