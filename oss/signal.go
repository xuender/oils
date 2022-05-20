package oss

import (
	"context"
	"os/signal"
	"syscall"
)

type Closer interface {
	Close() error
}

type Errorer interface {
	Error(error)
}

func SignalClose(errorer Errorer, closers ...Closer) context.Context {
	ctx, stop := signal.NotifyContext(context.Background(),
		// 挂起 终端连接断开
		syscall.SIGHUP,
		// 中断 Ctrl+C
		syscall.SIGINT,
		// 结束
		syscall.SIGTERM,
		// 退出 (Ctrl+/)
		syscall.SIGQUIT,
		// 杀死
		syscall.SIGKILL,
	)

	go func() {
		<-ctx.Done()

		for _, closer := range closers {
			if err := closer.Close(); err != nil {
				errorer.Error(err)
			}
		}

		stop()
	}()

	return ctx
}
