package times

import (
	"sync"
	"time"
)

// Debounced 防抖函数.
//
// Deprecated: Use github.com/samber/lo.NewDebounce instead.
func Debounced(fun func(), interval time.Duration) func() {
	var (
		mutex sync.Mutex
		timer *time.Timer
	)

	return func() {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(interval, fun)
	}
}
