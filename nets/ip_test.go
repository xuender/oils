package nets_test

import (
	"net"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/nets"
)

func TestGetIP(t *testing.T) {
	t.Parallel()

	localAddr := nets.GetIP()

	assert.Equal(t, 4, len(localAddr))
	assert.NotEqual(t, 0, localAddr[0])

	gomonkey.ApplyFuncReturn(net.Dial, nil, net.ErrClosed)

	localAddr = nets.GetIP()

	assert.Equal(t, 0, localAddr[0])
}
