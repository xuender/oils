package rdb_test

import (
	"testing"

	"github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/rdb"
	rdb_mock "github.com/xuender/oils/syncs/rdb/mock"
)

func TestNewLimit(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)

	assert.NotNil(t, rdb.NewLimiter(client, "test", 10))
}

func TestLimit_Wait1(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limit := rdb.NewLimiter(client, "test", 3)
	cmd1 := new(redis.BoolCmd)
	cmd2 := new(redis.IntCmd)
	cmd3 := new(redis.IntCmd)

	cmd1.SetVal(true)
	cmd2.SetVal(0)
	cmd3.SetVal(1)

	client.EXPECT().SetNX(gomock.Any(), "_test_", gomock.Any(), gomock.Any()).Return(cmd1).MinTimes(0).MaxTimes(1)
	client.EXPECT().Del(gomock.Any(), "test").Return(cmd2).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd3).MinTimes(0).MaxTimes(1)

	limit.Wait()
}

func TestLimit_Wait2(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limit := rdb.NewLimiter(client, "test", 10)
	cmd3 := new(redis.IntCmd)
	cmd4 := new(redis.BoolCmd)
	cmd5 := new(redis.IntCmd)
	cmd6 := new(redis.DurationCmd)

	cmd3.SetVal(1)
	cmd4.SetVal(false)
	cmd5.SetVal(2)
	cmd6.SetVal(2)

	client.EXPECT().SetNX(gomock.Any(), "_test_", gomock.Any(), gomock.Any()).Return(cmd4).MaxTimes(2)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd5).MaxTimes(1)
	client.EXPECT().PTTL(gomock.Any(), "_test_").Return(cmd6).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd3).MaxTimes(1)

	limit.Wait()
}

func TestLimit_Wait3(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limit := rdb.NewLimiter(client, "test", 10)
	cmd2 := new(redis.IntCmd)
	cmd3 := new(redis.IntCmd)
	cmd4 := new(redis.BoolCmd)
	cmd5 := new(redis.IntCmd)
	cmd6 := new(redis.DurationCmd)

	cmd2.SetVal(0)
	cmd3.SetVal(1)
	cmd4.SetVal(false)
	cmd5.SetVal(2)
	cmd6.SetVal(-1)

	client.EXPECT().SetNX(gomock.Any(), "_test_", gomock.Any(), gomock.Any()).Return(cmd4).MaxTimes(2)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd5).MaxTimes(1)
	client.EXPECT().PTTL(gomock.Any(), "_test_").Return(cmd6).MaxTimes(1)
	client.EXPECT().Del(gomock.Any(), "_test_").Return(cmd2).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd3).MaxTimes(1)

	limit.Wait()
}

func TestLimit_Try1(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limit := rdb.NewLimiter(client, "test", 3)
	cmd1 := new(redis.BoolCmd)
	cmd2 := new(redis.IntCmd)
	cmd3 := new(redis.IntCmd)

	cmd1.SetVal(true)
	cmd2.SetVal(0)
	cmd3.SetVal(1)

	client.EXPECT().SetNX(gomock.Any(), "_test_", gomock.Any(), gomock.Any()).Return(cmd1).MinTimes(0).MaxTimes(1)
	client.EXPECT().Del(gomock.Any(), "test").Return(cmd2).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd3).MinTimes(0).MaxTimes(1)

	_ = limit.Try()
}

func TestLimit_Try2(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	limit := rdb.NewLimiter(client, "test", 10)
	cmd3 := new(redis.IntCmd)
	cmd4 := new(redis.BoolCmd)
	cmd5 := new(redis.IntCmd)
	cmd6 := new(redis.DurationCmd)

	cmd3.SetVal(1)
	cmd4.SetVal(false)
	cmd5.SetVal(2)
	cmd6.SetVal(2)

	client.EXPECT().SetNX(gomock.Any(), "_test_", gomock.Any(), gomock.Any()).Return(cmd4).MaxTimes(2)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd5).MaxTimes(1)
	client.EXPECT().PTTL(gomock.Any(), "_test_").Return(cmd6).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "test").Return(cmd3).MaxTimes(1)

	_ = limit.Try()
}
