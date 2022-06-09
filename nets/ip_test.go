package nets_test

import (
	"fmt"
	"testing"

	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/nets"
)

func ExampleGetIP() {
	ip := nets.GetIP()
	fmt.Println(ip)
}

func TestGetIP(t *testing.T) {
	t.Parallel()

	ip := nets.GetIP()

	assert.Equal(t, 4, len(ip))
}
