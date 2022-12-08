package modes

import "sync"

// Subject 主题.
type Subject[DATA any, ID comparable] struct {
	observers map[ID]func(DATA)
	mutex     sync.RWMutex
}

func NewSubject[DATA any, ID comparable]() *Subject[DATA, ID] {
	return &Subject[DATA, ID]{
		observers: map[ID]func(DATA){},
		mutex:     sync.RWMutex{},
	}
}

// Register 注册.
func (p *Subject[DATA, ID]) Register(observerID ID, observer func(DATA)) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.observers[observerID] = observer
}

// Deregister 取消注册者.
func (p *Subject[DATA, ID]) Deregister(observerID ID) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	delete(p.observers, observerID)
}

// NotifyAll 通知所有.
func (p *Subject[DATA, ID]) NotifyAll(obj DATA) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, observer := range p.observers {
		observer(obj)
	}
}

// Notify 通知.
func (p *Subject[DATA, ID]) Notify(obj DATA, ids ...ID) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, id := range ids {
		if observer, has := p.observers[id]; has {
			observer(obj)
		}
	}
}
