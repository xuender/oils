package ios

type PassWriter struct{}

func NewPassWriter() *PassWriter {
	return &PassWriter{}
}

func (*PassWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
