package ios

import (
	"bytes"
	"io"
)

type ContainsReader struct {
	reader   io.Reader
	contains bool
	subslice []byte
}

func NewContainsReader(reader io.Reader, subslice []byte) *ContainsReader {
	return &ContainsReader{reader: reader, subslice: subslice}
}

func (p *ContainsReader) Read(data []byte) (n int, err error) {
	num, err := p.reader.Read(data)

	if !p.contains {
		p.contains = bytes.Contains(data, p.subslice)
	}

	return num, err
}

func (p *ContainsReader) Contains() bool {
	return p.contains
}
