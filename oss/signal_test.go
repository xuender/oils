package oss_test

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/oss"
)

type closer struct{ close bool }

func (p *closer) Close() error {
	p.close = true

	return os.ErrClosed
}

func (p *closer) Error(error) {}

// nolint: exhaustruct
func TestSignalClose(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	defer signal.Stop(c)

	clo := &closer{}
	oss.SignalClose(clo, clo)
	assert.False(t, clo.close)

	_ = syscall.Kill(syscall.Getpid(), syscall.SIGHUP)

	time.Sleep(time.Millisecond)
	assert.True(t, clo.close)
}
