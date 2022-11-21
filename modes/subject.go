package modes

// Subject 主题.
type Subject[OBJ any, ID comparable] struct {
	observers []Observer[OBJ, ID]
}

// Register 注册.
func (p *Subject[OBJ, ID]) Register(observer Observer[OBJ, ID]) {
	p.observers = append(p.observers, observer)
}

// Deregister 取消注册者.
func (p *Subject[OBJ, ID]) Deregister(observer Observer[OBJ, ID]) {
	p.observers = removeObserver(p.observers, observer)
}

// NotifyAll 通知所有.
func (p *Subject[OBJ, ID]) NotifyAll(obj OBJ) {
	for _, observer := range p.observers {
		observer.Update(obj)
	}
}

// Notify 通知.
func (p *Subject[OBJ, ID]) Notify(obj OBJ, ids ...ID) {
	for _, id := range ids {
		for _, observer := range p.observers {
			if observer.ID() == id {
				observer.Update(obj)
			}
		}
	}
}

func removeObserver[OBJ any, ID comparable](observers []Observer[OBJ, ID], elem Observer[OBJ, ID]) []Observer[OBJ, ID] {
	length := len(observers)

	for i, observer := range observers {
		if elem.ID() == observer.ID() {
			observers[length-1], observers[i] = observers[i], observers[length-1]

			return observers[:length-1]
		}
	}

	return observers
}
