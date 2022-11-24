package logs_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/logs"
)

// nolint: paralleltest
func TestInfo(t *testing.T) {
	type args struct {
		funcName string
		yield    func(...interface{})
	}

	for _, arg := range []args{
		{funcName: "Info", yield: logs.Info},
		{funcName: "Debug", yield: logs.Debug},
		{funcName: "Warn", yield: logs.Warn},
		{funcName: "Error", yield: logs.Error},
		{funcName: "DPanic", yield: logs.DPanic},
		{funcName: "Panic", yield: logs.Panic},
		{funcName: "Fatal", yield: logs.Fatal},
	} {
		data := 0
		patches := gomonkey.ApplyMethodFunc(logs.GetLog(), arg.funcName, func(args ...interface{}) {
			// nolint
			switch arg := args[0].(type) {
			case int:
				data = arg
			default:
			}
		})

		defer patches.Reset()

		arg.yield(1)
		assert.Equal(t, 1, data)
	}
}

// nolint: paralleltest
func TestInfof(t *testing.T) {
	type args struct {
		funcName string
		yield    func(string, ...interface{})
	}

	for _, arg := range []args{
		{funcName: "Infof", yield: logs.Infof},
		{funcName: "Debugf", yield: logs.Debugf},
		{funcName: "Warnf", yield: logs.Warnf},
		{funcName: "Errorf", yield: logs.Errorf},
		{funcName: "DPanicf", yield: logs.DPanicf},
		{funcName: "Panicf", yield: logs.Panicf},
		{funcName: "Fatalf", yield: logs.Fatalf},
		{funcName: "Infow", yield: logs.Infow},
		{funcName: "Debugw", yield: logs.Debugw},
		{funcName: "Warnw", yield: logs.Warnw},
		{funcName: "Errorw", yield: logs.Errorw},
		{funcName: "DPanicw", yield: logs.DPanicw},
		{funcName: "Panicw", yield: logs.Panicw},
		{funcName: "Fatalw", yield: logs.Fatalw},
	} {
		data := 0
		patches := gomonkey.ApplyMethodFunc(logs.GetLog(), arg.funcName, func(str string, args ...interface{}) {
			switch arg := args[0].(type) {
			case int:
				data = arg
			default:
			}
		})

		defer patches.Reset()

		arg.yield("", 1)
		assert.Equal(t, 1, data)
	}
}

// nolint: paralleltest
func TestSync(t *testing.T) {
	run := false
	patches := gomonkey.ApplyMethodFunc(logs.GetLog(), "Sync", func() error {
		run = true

		return nil
	})

	defer patches.Reset()
	logs.Sync()

	assert.True(t, run)
}

func TestSetInfoLevel(t *testing.T) {
	t.Parallel()

	logs.SetInfoLevel()
}
