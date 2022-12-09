package modes

import (
	"sync"
	"time"

	"github.com/xuender/oils/base"
)

type DelayEventBus[DATA any, KEY comparable] struct {
	EventBus[DATA, KEY]
	data     map[KEY]*Event[DATA, KEY]
	interval time.Duration
	lock     sync.Mutex
	cache    chan *Event[DATA, KEY]
}

func NewDelayEventBus[DATA any, KEY comparable](interval time.Duration) *DelayEventBus[DATA, KEY] {
	return NewDelayEventBusBySize[DATA, KEY](interval, base.Hundred)
}

func NewDelayEventBusBySize[DATA any, KEY comparable](interval time.Duration, size int) *DelayEventBus[DATA, KEY] {
	res := &DelayEventBus[DATA, KEY]{
		EventBus: EventBus[DATA, KEY]{
			observers: map[KEY][]chan<- DATA{},
			mutex:     sync.RWMutex{},
		},
		data:     map[KEY]*Event[DATA, KEY]{},
		interval: interval,
		lock:     sync.Mutex{},
		cache:    make(chan *Event[DATA, KEY], size),
	}

	go res.post()

	return res
}

func (p *DelayEventBus[DATA, KEY]) post() {
	ticker := time.NewTicker(p.interval)

	for {
		select {
		case <-ticker.C:
			p.lock.Lock()

			for key, event := range p.data {
				p.EventBus.Post(event)
				delete(p.data, key)
			}

			p.lock.Unlock()
		case event := <-p.cache:
			p.lock.Lock()
			p.data[event.Key] = event
			p.lock.Unlock()
		}
	}
}

func (p *DelayEventBus[DATA, KEY]) Post(event *Event[DATA, KEY]) {
	p.cache <- event
}
