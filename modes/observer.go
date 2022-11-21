package modes

// Observer 观察者.
type Observer[OBJ any, ID comparable] interface {
	Update(data OBJ)
	ID() ID
}
