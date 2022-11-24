package rdb_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/syncs/rdb"
	rdb_mock "github.com/xuender/oils/syncs/rdb/mock"
)

func TestNewLocker(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmd := rdb_mock.NewMockCmdable(ctrl)

	assert.NotNil(t, rdb.NewLocker(cmd, time.Second, time.Millisecond))
	assert.Panics(t, func() {
		rdb.NewLocker(cmd, time.Second, time.Minute)
	})
}

func TestLocker_Lock(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	cmd := rdb_mock.NewMockCmdable(ctrl)
	cmd1 := &redis.BoolCmd{}
	cmd2 := &redis.BoolCmd{}

	cmd1.SetVal(true)
	cmd2.SetVal(false)

	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd1).MinTimes(0).MaxTimes(1)
	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd2).MinTimes(1).MaxTimes(100)
	cmd.EXPECT().Del(gomock.Any(), "oils-test-LOCK").Return(redis.NewIntCmd(context.Background())).AnyTimes()
	cmd.EXPECT().Expire(gomock.Any(), "oils-test-LOCK", gomock.Any()).Return(cmd1).AnyTimes()

	locker := rdb.NewLocker(cmd, time.Second*2, time.Millisecond*500)
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
	cmd := rdb_mock.NewMockCmdable(ctrl)

	cmd1 := &redis.BoolCmd{}
	cmd1.SetVal(false)
	cmd1.SetErr(rdb.ErrLock)

	cmd.EXPECT().SetNX(gomock.Any(), "oils-test-LOCK", gomock.Any(), gomock.Any()).Return(cmd1).AnyTimes()

	locker := rdb.NewLocker(cmd)
	count := 0

	err := locker.Lock("test", func() error {
		count++
		time.Sleep(time.Second)

		return nil
	})

	assert.True(t, errors.Is(err, rdb.ErrLock))
}
