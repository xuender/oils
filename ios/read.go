package ios

import (
	"bufio"
	"io"
	"os"
)

const (
	defaultBufSize = 1024 * 1000
)

func ReaderLine(reader io.Reader, call func(line string) error) error {
	buf := bufio.NewReaderSize(reader, defaultBufSize)

	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if err := call(string(line)); err != nil {
			return err
		}
	}

	return nil
}

func ReadLine(file string, call func(line string) error) error {
	reader, err := os.Open(file)
	if err != nil {
		return err
	}
	defer reader.Close()

	return ReaderLine(reader, call)
}
