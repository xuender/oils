package main

import (
	"github.com/xuender/oils/logs"
)

func main() {
	logs.Info("test")
	logs.RotateJSONLog("/tmp")
	logs.Infow("info", "name", "tom")
	logs.Debugw("debug", "name", "tom")
}
