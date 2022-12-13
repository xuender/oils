package modes

type Register[DATA any, KEY comparable] interface {
	Regist(chan<- DATA, []KEY)
}
