package modes

import (
	"sync"

	"github.com/samber/lo"
	"github.com/xuender/oils/ios"
)

// Subject 主题.
type Subject[DATA any] struct {
	observers []chan<- DATA
	mutex     sync.RWMutex
}

func NewSubject[DATA any]() *Subject[DATA] {
	return &Subject[DATA]{
		observers: []chan<- DATA{},
		mutex:     sync.RWMutex{},
	}
}

// Register 监听者注册.
func (p *Subject[DATA]) Register(observer chan<- DATA) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.observers = append(p.observers, observer)
}

// Notify 通知所有监听者.
func (p *Subject[DATA]) Notify(data DATA) {
	p.mutex.RLock()

	closes := []int{}

	for index, observer := range p.observers {
		if ios.Write(data, observer) != nil {
			closes = append(closes, index)
		}
	}

	p.mutex.RUnlock()

	if len(closes) == 0 {
		return
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.observers = lo.Filter(p.observers, func(_ chan<- DATA, index int) bool {
		return !lo.Contains(closes, index)
	})
}
