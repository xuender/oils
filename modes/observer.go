package modes

// Observer 观察者.
type Observer[T any, I comparable] interface {
	Update(data T)
	ID() I
}
