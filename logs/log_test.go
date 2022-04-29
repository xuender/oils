package logs_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/logs"
)

func TestLogName(t *testing.T) {
	t.Parallel()

	switch runtime.GOOS {
	case "window":
		assert.Equal(t, filepath.Join(filepath.Base(os.Args[0]), "logs.log"), logs.LogName())
	default:
		assert.Equal(t, "/var/tmp/logs.log", logs.LogName())
	}
}
