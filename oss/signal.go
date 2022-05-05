package oss

import (
	"os"
	"os/signal"
	"syscall"
)

type Closer interface {
	Close() error
}

type Errorer interface {
	Error(error)
}

func Signal() {
	sigChan := make(chan os.Signal, 1)
	// 监听指定信号 ctrl+c kill等
	signal.Notify(sigChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigChan
}

func SignalClose(errorer Errorer, closers ...Closer) {
	if len(closers) == 0 {
		return
	}

	sigChan := make(chan os.Signal, 1)
	// 监听指定信号 ctrl+c kill等
	signal.Notify(sigChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan

		for _, closer := range closers {
			if err := closer.Close(); err != nil {
				errorer.Error(err)
			}
		}
	}()
}
