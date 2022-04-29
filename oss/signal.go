package oss

import (
	"os"
	"os/signal"
	"syscall"
)

type closer interface {
	Close() error
}

type errorer interface {
	Error(error)
}

func SignalClose(call errorer, closers ...closer) {
	if len(closers) == 0 {
		return
	}

	sigChan := make(chan os.Signal, 1)
	// 监听指定信号 ctrl+c kill等
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan

		for _, closer := range closers {
			if err := closer.Close(); err != nil {
				call.Error(err)
			}
		}
	}()
}
