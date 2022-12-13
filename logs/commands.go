package logs

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	logrus.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	logrus.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logrus.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logrus.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logrus.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	logrus.Fatalf(template, args...)
}

func fields(keysAndValues []any) logrus.Fields {
	if len(keysAndValues)%2 != 0 {
		keysAndValues = append(keysAndValues, "")
	}

	fields := logrus.Fields{}

	for i := 0; i < len(keysAndValues); i += 2 {
		fields[fmt.Sprintf("%v", keysAndValues[i])] = keysAndValues[i+1]
	}

	return fields
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Debug(msg)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Info(msg)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Warn(msg)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Error(msg)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Panic(msg)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logrus.WithFields(fields(keysAndValues)).Fatal(msg)
}

func SetInfoLevel() {
	logrus.SetLevel(logrus.InfoLevel)
}
