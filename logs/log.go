package logs

import (
	"os"
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// nolint
var log = NewDebug().WithOptions(zap.AddCallerSkip(1)).Sugar()

func RotateLog(paths ...string) {
	week := 7
	filename := LogName(paths...)

	log.Info(filename)

	log = NewRotate(&lumberjack.Logger{
		Filename: filename,
		Compress: true,
		MaxAge:   week,
	}).WithOptions(zap.AddCallerSkip(1)).Sugar()
}

func NewDebug() *zap.Logger {
	return newLogger(zap.DebugLevel)
}

func NewInfo() *zap.Logger {
	return newLogger(zap.InfoLevel)
}

func newLogger(level zapcore.Level) *zap.Logger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      level == zap.DebugLevel,
		Encoding:         "console",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	if config.Development {
		config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05.00000")
	} else {
		config.EncoderConfig = zap.NewProductionEncoderConfig()
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("04:05.00000")
		config.EncoderConfig.EncodeCaller = nil
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	return logger
}

func LogName(paths ...string) string {
	command := os.Args[0]
	name := filepath.Base(command)
	ext := filepath.Ext(name)
	name = name[:len(name)-len(ext)] + ".log"

	if len(paths) > 0 {
		return filepath.Join(filepath.Join(paths...), name)
	}

	if runtime.GOOS == "windows" {
		return filepath.Join(filepath.Dir(command), name)
	}

	return filepath.Join("/var/tmp", name)
}
