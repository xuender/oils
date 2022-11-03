package nets_test

import (
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/nets"
)

func TestGetIP(t *testing.T) {
	t.Parallel()

	ip := nets.GetIP()

	assert.Equal(t, 4, len(ip))
}
