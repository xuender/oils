package ios_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/ios"
)

func TestNewWriter(t *testing.T) {
	t.Parallel()

	file := filepath.Join(os.TempDir(), "test_file")
	writer := base.Must1(ios.NewWriter(file))

	_, _ = writer.Write([]byte("1"))
	writer.Close()

	writer = base.Must1(ios.NewWriter(file))
	_, _ = writer.Write([]byte("2"))
	writer.Close()

	data := base.Must1(os.ReadFile(file))

	assert.Equal(t, "12", string(data))
	os.Remove(file)
}

func TestNewWriter_error(t *testing.T) {
	t.Parallel()

	_, err := ios.NewWriter(os.TempDir())
	assert.NotNil(t, err)
}
