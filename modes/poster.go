package modes

type Poster[DATA any, KEY comparable] interface {
	Post(*Event[DATA, KEY])
}
