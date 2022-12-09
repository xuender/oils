package modes

type Event[DATA any, KEY comparable] struct {
	Key  KEY
	Data DATA
}

func NewEvent[DATA any, KEY comparable](key KEY, data DATA) *Event[DATA, KEY] {
	return &Event[DATA, KEY]{
		Key:  key,
		Data: data,
	}
}
