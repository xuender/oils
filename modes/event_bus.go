package modes

import (
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/xuender/oils/ios"
)

// EventBus 消息总线.
type EventBus[DATA any, KEY comparable] struct {
	observers map[KEY][]chan<- DATA
	mutex     sync.RWMutex
}

// NewEventBus 新建消息总线.
func NewEventBus[DATA any, KEY comparable]() *EventBus[DATA, KEY] {
	return &EventBus[DATA, KEY]{
		observers: map[KEY][]chan<- DATA{},
		mutex:     sync.RWMutex{},
	}
}

// Regist 注册观察者.
func (p *EventBus[DATA, KEY]) Regist(objerver chan<- DATA, keys []KEY) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, key := range keys {
		if ids, has := p.observers[key]; has {
			p.observers[key] = append(ids, objerver)
		} else {
			p.observers[key] = []chan<- DATA{objerver}
		}
	}
}

// Post 发送消息.
func (p *EventBus[DATA, KEY]) Post(event *Event[DATA, KEY]) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	obs, has := p.observers[event.Key]
	if !has {
		return
	}

	closes := []int{}

	for index, observer := range obs {
		if err := ios.WriteTimeout(event.Data, observer, time.Second); err != nil {
			closes = append(closes, index)
		}
	}

	observers := lo.Filter(obs, func(_ chan<- DATA, index int) bool {
		return !lo.Contains(closes, index)
	})

	if len(observers) == 0 {
		delete(p.observers, event.Key)

		return
	}

	p.observers[event.Key] = observers
}

// Has 消息类型是否有监听.
func (p *EventBus[DATA, KEY]) Has(key KEY) bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	_, has := p.observers[key]

	return has
}
