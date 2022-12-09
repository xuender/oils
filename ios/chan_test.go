package ios_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/ios"
)

func TestInput(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	data := make(chan int, 2)

	assert.Nil(ios.Write(1, data))
	close(data)
	assert.NotNil(ios.Write(1, data))
}

func TestWriteTimeout(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	data := make(chan int, 1)

	assert.Nil(ios.WriteTimeout(1, data, time.Second))
	assert.NotNil(ios.WriteTimeout(1, data, time.Millisecond))
	<-data
	close(data)
	assert.NotNil(ios.WriteTimeout(1, data, time.Second))
}

func TestRead(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	data := make(chan int, 1)

	data <- 1

	res, err := ios.Read(data)

	assert.Nil(err)
	assert.Equal(1, res)

	close(data)

	_, err = ios.Read(data)

	assert.NotNil(err)
}

func TestReadTimeout(t *testing.T) {
	t.Parallel()

	assert := assert.New(t)
	data := make(chan int, 1)

	data <- 1

	num, err := ios.ReadTimeout(data, time.Second)

	assert.Nil(err)
	assert.Equal(1, num)

	_, err = ios.ReadTimeout(data, time.Millisecond)
	assert.NotNil(err)

	data <- 2

	close(data)

	num, err = ios.ReadTimeout(data, time.Millisecond)

	assert.Equal(2, num)
	assert.Nil(err)

	_, err = ios.ReadTimeout(data, time.Millisecond)

	assert.NotNil(err)
}
