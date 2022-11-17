package modes

// Subject 主题.
type Subject[T any, I comparable] struct {
	observers []Observer[T, I]
}

// Register 注册.
func (p *Subject[T, I]) Register(o Observer[T, I]) {
	p.observers = append(p.observers, o)
}

// Deregister 取消注册者.
func (p *Subject[T, I]) Deregister(o Observer[T, I]) {
	p.observers = removeObserver(p.observers, o)
}

// NotifyAll 通知所有.
func (p *Subject[T, I]) NotifyAll(obj T) {
	for _, observer := range p.observers {
		observer.Update(obj)
	}
}

// Notify 通知.
func (p *Subject[T, I]) Notify(obj T, ids ...I) {
	for _, id := range ids {
		for _, observer := range p.observers {
			if observer.ID() == id {
				observer.Update(obj)
			}
		}
	}
}

func removeObserver[T any, I comparable](observers []Observer[T, I], elem Observer[T, I]) []Observer[T, I] {
	length := len(observers)

	for i, observer := range observers {
		if elem.ID() == observer.ID() {
			observers[length-1], observers[i] = observers[i], observers[length-1]

			return observers[:length-1]
		}
	}

	return observers
}
