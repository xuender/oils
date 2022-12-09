package nets_test

import (
	"net"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/nets"
)

// nolint: paralleltest
func TestGetIP(t *testing.T) {
	localAddr := nets.GetIP()

	assert.Equal(t, 4, len(localAddr))
	assert.NotEqual(t, 0, localAddr[0])

	patch := gomonkey.ApplyFuncReturn(net.Dial, nil, net.ErrClosed)
	defer patch.Reset()

	localAddr = nets.GetIP()

	assert.Equal(t, byte(0), localAddr[0])
}
