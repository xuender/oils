package ios

import (
	"time"
)

// Write 返回管道写入是否失败.
func Write[E any](data E, target chan<- E) (err error) {
	defer func() {
		if recover() != nil {
			err = ErrChanClose
		}
	}()

	target <- data

	return
}

// WriteTimeout 返回管道写入是否失败.
func WriteTimeout[E any](data E, target chan<- E, timeout time.Duration) (err error) {
	defer func() {
		if recover() != nil {
			err = ErrChanClose
		}
	}()

	select {
	case target <- data:
	case <-time.After(timeout):
		return ErrTimeout
	}

	return err
}

// Read 读取管道.
func Read[E any](source <-chan E) (E, error) {
	data, has := <-source

	if has {
		return data, nil
	}

	return data, ErrChanClose
}

// Read 限制时间内读取管道.
func ReadTimeout[E any](source <-chan E, timeout time.Duration) (E, error) {
	select {
	case data, has := <-source:
		if has {
			return data, nil
		}

		return data, ErrChanClose
	case <-time.After(timeout):
		var zero E

		return zero, ErrTimeout
	}
}
