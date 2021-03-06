package logs_test

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/logs"
)

func TestInfo(t *testing.T) {
	t.Parallel()

	type args struct {
		funcName string
		funcCall func(...interface{})
	}

	for _, arg := range []args{
		{funcName: "Info", funcCall: logs.Info},
		{funcName: "Debug", funcCall: logs.Debug},
		{funcName: "Warn", funcCall: logs.Warn},
		{funcName: "Error", funcCall: logs.Error},
		{funcName: "DPanic", funcCall: logs.DPanic},
		{funcName: "Panic", funcCall: logs.Panic},
		{funcName: "Fatal", funcCall: logs.Fatal},
	} {
		data := 0
		patches := gomonkey.ApplyMethodFunc(logs.GetLog(), arg.funcName, func(args interface{}) {
			// nolint
			switch arg := args.(type) {
			case []interface{}:
				switch num := arg[0].(type) {
				case int:
					data = num
				}
			}
		})

		defer patches.Reset()

		arg.funcCall(1)
		assert.Equal(t, 1, data)
	}
}

func TestInfof(t *testing.T) {
	t.Parallel()

	type args struct {
		funcName string
		funcCall func(string, ...interface{})
	}

	for _, arg := range []args{
		{funcName: "Infof", funcCall: logs.Infof},
		{funcName: "Debugf", funcCall: logs.Debugf},
		{funcName: "Warnf", funcCall: logs.Warnf},
		{funcName: "Errorf", funcCall: logs.Errorf},
		{funcName: "DPanicf", funcCall: logs.DPanicf},
		{funcName: "Panicf", funcCall: logs.Panicf},
		{funcName: "Fatalf", funcCall: logs.Fatalf},
		{funcName: "Infow", funcCall: logs.Infow},
		{funcName: "Debugw", funcCall: logs.Debugw},
		{funcName: "Warnw", funcCall: logs.Warnw},
		{funcName: "Errorw", funcCall: logs.Errorw},
		{funcName: "DPanicw", funcCall: logs.DPanicw},
		{funcName: "Panicw", funcCall: logs.Panicw},
		{funcName: "Fatalw", funcCall: logs.Fatalw},
	} {
		data := 0
		patches := gomonkey.ApplyMethodFunc(logs.GetLog(), arg.funcName, func(str string, args ...interface{}) {
			// nolint
			switch arg := args[0].(type) {
			case []interface{}:
				switch num := arg[0].(type) {
				case int:
					data = num
				}
			}
		})

		defer patches.Reset()

		arg.funcCall("", 1)
		assert.Equal(t, 1, data)
	}
}
