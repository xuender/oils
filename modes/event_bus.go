package modes

import (
	"sync"

	"github.com/samber/lo"
	"github.com/xuender/oils/ios"
)

type EventBus[DATA any, KEY comparable] struct {
	observers map[KEY][]chan<- DATA
	mutex     sync.RWMutex
}

func NewEventBus[DATA any, KEY comparable]() *EventBus[DATA, KEY] {
	return &EventBus[DATA, KEY]{
		observers: map[KEY][]chan<- DATA{},
		mutex:     sync.RWMutex{},
	}
}

func (p *EventBus[DATA, KEY]) Register(objerver chan<- DATA, keys []KEY) {
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

func (p *EventBus[DATA, KEY]) Post(event *Event[DATA, KEY]) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	obs, has := p.observers[event.Key]
	if !has {
		return
	}

	closes := []int{}

	for index, observer := range obs {
		if err := ios.Write(event.Data, observer); err != nil {
			closes = append(closes, index)
		}
	}

	p.observers[event.Key] = lo.Filter(obs, func(_ chan<- DATA, index int) bool {
		return !lo.Contains(closes, index)
	})
}
