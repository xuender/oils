package times

import (
	"sync"
	"time"
)

// Debounced 防抖函数.
func Debounced(fun func(), interval time.Duration) func() {
	var mutex sync.Mutex
	var timer *time.Timer

	return func() {
		mutex.Lock()
		defer mutex.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(interval, fun)
	}
}
