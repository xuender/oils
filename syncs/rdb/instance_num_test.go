package rdb_test

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
	"github.com/xuender/oils/assert"
	"github.com/xuender/oils/syncs/rdb"
	rdb_mock "github.com/xuender/oils/syncs/rdb/mock"
)

type Customer struct {
	id int
}

func (p *Customer) ID() int {
	return p.id
}

func (p *Customer) Update(data uint64) {
	// nolint
	fmt.Println(p.id, data)
}

func TestInstanceNum(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	ins := rdb.NewInstanceNum(client)
	update := make(chan uint64)

	cmd1 := new(redis.IntCmd)
	cmd2 := new(redis.IntCmd)

	cmd1.SetVal(1)
	cmd2.SetVal(2)
	client.EXPECT().Incr(gomock.Any(), "ins_num").Return(cmd1).MinTimes(0).MaxTimes(1)
	client.EXPECT().Incr(gomock.Any(), "ins_num").Return(cmd2).MinTimes(1).MaxTimes(20)

	assert.Equal(t, 1, ins.Num())

	go ins.Run()

	time.Sleep(time.Second * 2)
	assert.Equal(t, 1, ins.Num())
	ins.Register(update)
}

func TestInstanceNumError(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	client := rdb_mock.NewMockCmdable(ctrl)
	ins := rdb.NewInstanceNum(client)

	cmd1 := new(redis.IntCmd)
	cmd2 := new(redis.IntCmd)

	cmd1.SetErr(io.ErrClosedPipe)
	client.EXPECT().Incr(gomock.Any(), "ins_num").Return(cmd1).MinTimes(0).MaxTimes(1)
	client.EXPECT().Del(gomock.Any(), "ins_num").Return(cmd1).MinTimes(0).MaxTimes(1)

	cmd2.SetVal(1)
	client.EXPECT().Incr(gomock.Any(), "ins_num").Return(cmd2).MinTimes(1).MaxTimes(30)

	assert.Equal(t, 1, ins.Num())

	go ins.Run()

	time.Sleep(time.Second * 2)
}
