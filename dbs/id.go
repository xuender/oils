package dbs

import (
	"time"

	"github.com/xuender/oils/base"
	"github.com/xuender/oils/nets"
)

const (
	// min 2021-10-31.
	min            = 1635638400000
	tsPosition     = 12
	serialPosition = 4
	machineMax     = 16
	serialMax      = 256
)

// nolint
var (
	chanUint64 = make(chan uint64, 1)
	chanBool   = make(chan struct{}, 1)
	ipAddress  uint64
)

// NewID 新建ID.
// num 4位 16个.
// 起始日期 2021-10-31 日.
// 合计53位，js中使用不会出错.
func NewID(num uint64) uint64 {
	// 40 位作为毫秒数(35年)
	// 8 位序号(256个)
	// 4 位服务器代码(16个)
	return genID(num % machineMax)
}

// ID 获取ID.
func ID() uint64 {
	return genID(ipAddress)
}

// DecodeID id解码.
func DecodeID(id uint64) (uint64, uint64, uint64) {
	machine := id % machineMax
	serial := (id >> serialPosition) % serialMax
	ts := id >> tsPosition

	return ts, serial, machine
}

func genID(machine uint64) uint64 {
	chanBool <- base.None

	return machine | <-chanUint64
}

func getSerial() {
	var (
		serial uint8
		stamp  = time.Now().UnixMilli() - min
		now    int64
	)

	for {
		<-chanBool

		now = time.Now().UnixMilli() - min
		for ; serial == 0 && now == stamp; now = time.Now().UnixMilli() - min {
			// fmt.Println("sleep")
			time.Sleep(time.Microsecond)
		}

		chanUint64 <- uint64(stamp<<tsPosition) + (uint64(serial) << serialPosition)

		stamp = now
		serial++
	}
}

// nolint
func init() {
	ipAddress = uint64(nets.GetIP()[3] % machineMax)

	go getSerial()
}
