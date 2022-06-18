package distributed_test

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/distributed"
)

func TestNewLocker(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmd := NewMockCmdable(ctrl)

	assert.NotNil(t, distributed.NewLocker(cmd, time.Second, time.Millisecond))
}

func TestLocker_Lock(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmd := NewMockCmdable(ctrl)
	cmd1 := &redis.BoolCmd{}
	cmd2 := &redis.BoolCmd{}

	cmd1.SetVal(true)
	cmd2.SetVal(false)

	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd1).MinTimes(0).MaxTimes(1)
	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd2).MinTimes(1).MaxTimes(100)
	cmd.EXPECT().Del(gomock.Any(), "oils-test-LOCK").Return(redis.NewIntCmd(nil)).AnyTimes()
	cmd.EXPECT().Expire(gomock.Any(), "oils-test-LOCK", gomock.Any()).Return(cmd1).AnyTimes()

	locker := distributed.NewLocker(cmd)
	count := 0

	for i := 0; i < 10; i++ {
		go base.Pass(locker.Lock("test", func() error {
			count++
			time.Sleep(time.Second)

			return nil
		}))
	}

	time.Sleep(time.Second)
	assert.Equal(t, 1, count)
}

func TestLocker_Lock_Error(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmd := NewMockCmdable(ctrl)

	cmd1 := &redis.BoolCmd{}
	cmd1.SetVal(false)
	cmd1.SetErr(distributed.ErrLock)

	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd1).AnyTimes()

	locker := distributed.NewLocker(cmd)
	count := 0

	err := locker.Lock("test", func() error {
		count++
		time.Sleep(time.Second)

		return nil
	})

	assert.True(t, errors.Is(err, distributed.ErrLock))
}
