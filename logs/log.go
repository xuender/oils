package logs

import (
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/xuender/oils/ios"
	"github.com/xuender/oils/oss"
)

func getHook(filename string, formatter logrus.Formatter) logrus.Hook {
	writer, err := rotatelogs.New(
		filename+".%Y%m%d%H",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		// rotatelogs.WithRotationCount(100),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, formatter)
}

func RotateLog(paths ...string) {
	filename := LogName(paths...)

	logrus.SetOutput(ios.NewPassWriter())
	logrus.AddHook(getHook(filename, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"}))
}

func RotateJSONLog(paths ...string) {
	filename := LogName(paths...)

	logrus.SetOutput(ios.NewPassWriter())
	logrus.AddHook(getHook(filename, &logrus.JSONFormatter{}))
}

func LogName(paths ...string) string {
	command := os.Args[0]
	name := filepath.Base(command)
	ext := filepath.Ext(name)
	name = name[:len(name)-len(ext)] + ".log"

	if len(paths) > 0 {
		return filepath.Join(filepath.Join(paths...), name)
	}

	if oss.IsWindows() {
		return filepath.Join(filepath.Dir(command), name)
	}

	return filepath.Join("/var/tmp", name)
}
