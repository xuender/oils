package modes

import (
	"sync"
	"time"

	"github.com/xuender/oils/base"
)

// DelayEventBus 延迟消息总线.
type DelayEventBus[DATA any, KEY comparable] struct {
	EventBus[DATA, KEY]
	data     map[KEY]*Event[DATA, KEY]
	interval time.Duration
	lock     sync.Mutex
	cache    chan *Event[DATA, KEY]
}

// NewDelayEventBus 新建延迟消息总线.
func NewDelayEventBus[DATA any, KEY comparable](interval time.Duration) *DelayEventBus[DATA, KEY] {
	return NewDelayEventBusBySize[DATA, KEY](interval, base.Hundred)
}

// NewDelayEventBusBySize 根据缓存尺寸新建延迟消息总线.
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
			if len(p.data) == 0 {
				continue
			}

			p.lock.Lock()

			for key, event := range p.data {
				p.EventBus.Post(event)
				delete(p.data, key)
			}

			p.lock.Unlock()
		case event := <-p.cache:
			if p.EventBus.Has(event.Key) {
				p.lock.Lock()
				p.data[event.Key] = event
				p.lock.Unlock()
			}
		}
	}
}

// Post 发送消息.
func (p *DelayEventBus[DATA, KEY]) Post(event *Event[DATA, KEY]) {
	p.cache <- event
}
