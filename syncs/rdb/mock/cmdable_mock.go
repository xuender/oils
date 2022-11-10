// nolint
package rdb_mock

import (
	context "context"
	reflect "reflect"
	time "time"

	"github.com/go-redis/redis/v8"
	gomock "github.com/golang/mock/gomock"
)

// MockCmdable is a mock of Cmdable interface.
type MockCmdable struct {
	ctrl     *gomock.Controller
	recorder *MockCmdableMockRecorder
}

// MockCmdableMockRecorder is the mock recorder for MockCmdable.
type MockCmdableMockRecorder struct {
	mock *MockCmdable
}

// NewMockCmdable creates a new mock instance.
func NewMockCmdable(ctrl *gomock.Controller) *MockCmdable {
	mock := &MockCmdable{ctrl: ctrl}
	mock.recorder = &MockCmdableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCmdable) EXPECT() *MockCmdableMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockCmdable) Append(ctx context.Context, key, value string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", ctx, key, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Append indicates an expected call of Append.
func (mr *MockCmdableMockRecorder) Append(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockCmdable)(nil).Append), ctx, key, value)
}

// BLMove mocks base method.
func (m *MockCmdable) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BLMove", ctx, source, destination, srcpos, destpos, timeout)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BLMove indicates an expected call of BLMove.
func (mr *MockCmdableMockRecorder) BLMove(ctx, source, destination, srcpos, destpos, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLMove", reflect.TypeOf((*MockCmdable)(nil).BLMove), ctx, source, destination, srcpos, destpos, timeout)
}

// BLPop mocks base method.
func (m *MockCmdable) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BLPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BLPop indicates an expected call of BLPop.
func (mr *MockCmdableMockRecorder) BLPop(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLPop", reflect.TypeOf((*MockCmdable)(nil).BLPop), varargs...)
}

// BRPop mocks base method.
func (m *MockCmdable) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BRPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BRPop indicates an expected call of BRPop.
func (mr *MockCmdableMockRecorder) BRPop(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPop", reflect.TypeOf((*MockCmdable)(nil).BRPop), varargs...)
}

// BRPopLPush mocks base method.
func (m *MockCmdable) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BRPopLPush", ctx, source, destination, timeout)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// BRPopLPush indicates an expected call of BRPopLPush.
func (mr *MockCmdableMockRecorder) BRPopLPush(ctx, source, destination, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BRPopLPush", reflect.TypeOf((*MockCmdable)(nil).BRPopLPush), ctx, source, destination, timeout)
}

// BZPopMax mocks base method.
func (m *MockCmdable) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMax indicates an expected call of BZPopMax.
func (mr *MockCmdableMockRecorder) BZPopMax(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMax", reflect.TypeOf((*MockCmdable)(nil).BZPopMax), varargs...)
}

// BZPopMin mocks base method.
func (m *MockCmdable) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZWithKeyCmd)
	return ret0
}

// BZPopMin indicates an expected call of BZPopMin.
func (mr *MockCmdableMockRecorder) BZPopMin(ctx, timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BZPopMin", reflect.TypeOf((*MockCmdable)(nil).BZPopMin), varargs...)
}

// BgRewriteAOF mocks base method.
func (m *MockCmdable) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgRewriteAOF", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgRewriteAOF indicates an expected call of BgRewriteAOF.
func (mr *MockCmdableMockRecorder) BgRewriteAOF(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgRewriteAOF", reflect.TypeOf((*MockCmdable)(nil).BgRewriteAOF), ctx)
}

// BgSave mocks base method.
func (m *MockCmdable) BgSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BgSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// BgSave indicates an expected call of BgSave.
func (mr *MockCmdableMockRecorder) BgSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BgSave", reflect.TypeOf((*MockCmdable)(nil).BgSave), ctx)
}

// BitCount mocks base method.
func (m *MockCmdable) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitCount", ctx, key, bitCount)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitCount indicates an expected call of BitCount.
func (mr *MockCmdableMockRecorder) BitCount(ctx, key, bitCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitCount", reflect.TypeOf((*MockCmdable)(nil).BitCount), ctx, key, bitCount)
}

// BitField mocks base method.
func (m *MockCmdable) BitField(ctx context.Context, key string, args ...interface{}) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitField", varargs...)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// BitField indicates an expected call of BitField.
func (mr *MockCmdableMockRecorder) BitField(ctx, key interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitField", reflect.TypeOf((*MockCmdable)(nil).BitField), varargs...)
}

// BitOpAnd mocks base method.
func (m *MockCmdable) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpAnd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpAnd indicates an expected call of BitOpAnd.
func (mr *MockCmdableMockRecorder) BitOpAnd(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpAnd", reflect.TypeOf((*MockCmdable)(nil).BitOpAnd), varargs...)
}

// BitOpNot mocks base method.
func (m *MockCmdable) BitOpNot(ctx context.Context, destKey, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BitOpNot", ctx, destKey, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpNot indicates an expected call of BitOpNot.
func (mr *MockCmdableMockRecorder) BitOpNot(ctx, destKey, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpNot", reflect.TypeOf((*MockCmdable)(nil).BitOpNot), ctx, destKey, key)
}

// BitOpOr mocks base method.
func (m *MockCmdable) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpOr", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpOr indicates an expected call of BitOpOr.
func (mr *MockCmdableMockRecorder) BitOpOr(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpOr", reflect.TypeOf((*MockCmdable)(nil).BitOpOr), varargs...)
}

// BitOpXor mocks base method.
func (m *MockCmdable) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destKey}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitOpXor", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitOpXor indicates an expected call of BitOpXor.
func (mr *MockCmdableMockRecorder) BitOpXor(ctx, destKey interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destKey}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitOpXor", reflect.TypeOf((*MockCmdable)(nil).BitOpXor), varargs...)
}

// BitPos mocks base method.
func (m *MockCmdable) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, bit}
	for _, a := range pos {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BitPos", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// BitPos indicates an expected call of BitPos.
func (mr *MockCmdableMockRecorder) BitPos(ctx, key, bit interface{}, pos ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, bit}, pos...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BitPos", reflect.TypeOf((*MockCmdable)(nil).BitPos), varargs...)
}

// ClientGetName mocks base method.
func (m *MockCmdable) ClientGetName(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetName", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientGetName indicates an expected call of ClientGetName.
func (mr *MockCmdableMockRecorder) ClientGetName(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetName", reflect.TypeOf((*MockCmdable)(nil).ClientGetName), ctx)
}

// ClientID mocks base method.
func (m *MockCmdable) ClientID(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockCmdableMockRecorder) ClientID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockCmdable)(nil).ClientID), ctx)
}

// ClientKill mocks base method.
func (m *MockCmdable) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientKill", ctx, ipPort)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClientKill indicates an expected call of ClientKill.
func (mr *MockCmdableMockRecorder) ClientKill(ctx, ipPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKill", reflect.TypeOf((*MockCmdable)(nil).ClientKill), ctx, ipPort)
}

// ClientKillByFilter mocks base method.
func (m *MockCmdable) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClientKillByFilter", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClientKillByFilter indicates an expected call of ClientKillByFilter.
func (mr *MockCmdableMockRecorder) ClientKillByFilter(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientKillByFilter", reflect.TypeOf((*MockCmdable)(nil).ClientKillByFilter), varargs...)
}

// ClientList mocks base method.
func (m *MockCmdable) ClientList(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientList", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClientList indicates an expected call of ClientList.
func (mr *MockCmdableMockRecorder) ClientList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientList", reflect.TypeOf((*MockCmdable)(nil).ClientList), ctx)
}

// ClientPause mocks base method.
func (m *MockCmdable) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientPause", ctx, dur)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ClientPause indicates an expected call of ClientPause.
func (mr *MockCmdableMockRecorder) ClientPause(ctx, dur interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientPause", reflect.TypeOf((*MockCmdable)(nil).ClientPause), ctx, dur)
}

// ClusterAddSlots mocks base method.
func (m *MockCmdable) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range slots {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterAddSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlots indicates an expected call of ClusterAddSlots.
func (mr *MockCmdableMockRecorder) ClusterAddSlots(ctx interface{}, slots ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, slots...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterAddSlots), varargs...)
}

// ClusterAddSlotsRange mocks base method.
func (m *MockCmdable) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterAddSlotsRange", ctx, min, max)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterAddSlotsRange indicates an expected call of ClusterAddSlotsRange.
func (mr *MockCmdableMockRecorder) ClusterAddSlotsRange(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterAddSlotsRange", reflect.TypeOf((*MockCmdable)(nil).ClusterAddSlotsRange), ctx, min, max)
}

// ClusterCountFailureReports mocks base method.
func (m *MockCmdable) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountFailureReports", ctx, nodeID)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountFailureReports indicates an expected call of ClusterCountFailureReports.
func (mr *MockCmdableMockRecorder) ClusterCountFailureReports(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountFailureReports", reflect.TypeOf((*MockCmdable)(nil).ClusterCountFailureReports), ctx, nodeID)
}

// ClusterCountKeysInSlot mocks base method.
func (m *MockCmdable) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterCountKeysInSlot", ctx, slot)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterCountKeysInSlot indicates an expected call of ClusterCountKeysInSlot.
func (mr *MockCmdableMockRecorder) ClusterCountKeysInSlot(ctx, slot interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterCountKeysInSlot", reflect.TypeOf((*MockCmdable)(nil).ClusterCountKeysInSlot), ctx, slot)
}

// ClusterDelSlots mocks base method.
func (m *MockCmdable) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range slots {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ClusterDelSlots", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlots indicates an expected call of ClusterDelSlots.
func (mr *MockCmdableMockRecorder) ClusterDelSlots(ctx interface{}, slots ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, slots...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterDelSlots), varargs...)
}

// ClusterDelSlotsRange mocks base method.
func (m *MockCmdable) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterDelSlotsRange", ctx, min, max)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterDelSlotsRange indicates an expected call of ClusterDelSlotsRange.
func (mr *MockCmdableMockRecorder) ClusterDelSlotsRange(ctx, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterDelSlotsRange", reflect.TypeOf((*MockCmdable)(nil).ClusterDelSlotsRange), ctx, min, max)
}

// ClusterFailover mocks base method.
func (m *MockCmdable) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterFailover", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterFailover indicates an expected call of ClusterFailover.
func (mr *MockCmdableMockRecorder) ClusterFailover(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterFailover", reflect.TypeOf((*MockCmdable)(nil).ClusterFailover), ctx)
}

// ClusterForget mocks base method.
func (m *MockCmdable) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterForget", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterForget indicates an expected call of ClusterForget.
func (mr *MockCmdableMockRecorder) ClusterForget(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterForget", reflect.TypeOf((*MockCmdable)(nil).ClusterForget), ctx, nodeID)
}

// ClusterGetKeysInSlot mocks base method.
func (m *MockCmdable) ClusterGetKeysInSlot(ctx context.Context, slot, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterGetKeysInSlot", ctx, slot, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterGetKeysInSlot indicates an expected call of ClusterGetKeysInSlot.
func (mr *MockCmdableMockRecorder) ClusterGetKeysInSlot(ctx, slot, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterGetKeysInSlot", reflect.TypeOf((*MockCmdable)(nil).ClusterGetKeysInSlot), ctx, slot, count)
}

// ClusterInfo mocks base method.
func (m *MockCmdable) ClusterInfo(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterInfo", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterInfo indicates an expected call of ClusterInfo.
func (mr *MockCmdableMockRecorder) ClusterInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterInfo", reflect.TypeOf((*MockCmdable)(nil).ClusterInfo), ctx)
}

// ClusterKeySlot mocks base method.
func (m *MockCmdable) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterKeySlot", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ClusterKeySlot indicates an expected call of ClusterKeySlot.
func (mr *MockCmdableMockRecorder) ClusterKeySlot(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterKeySlot", reflect.TypeOf((*MockCmdable)(nil).ClusterKeySlot), ctx, key)
}

// ClusterMeet mocks base method.
func (m *MockCmdable) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterMeet", ctx, host, port)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterMeet indicates an expected call of ClusterMeet.
func (mr *MockCmdableMockRecorder) ClusterMeet(ctx, host, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterMeet", reflect.TypeOf((*MockCmdable)(nil).ClusterMeet), ctx, host, port)
}

// ClusterNodes mocks base method.
func (m *MockCmdable) ClusterNodes(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterNodes", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ClusterNodes indicates an expected call of ClusterNodes.
func (mr *MockCmdableMockRecorder) ClusterNodes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterNodes", reflect.TypeOf((*MockCmdable)(nil).ClusterNodes), ctx)
}

// ClusterReplicate mocks base method.
func (m *MockCmdable) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterReplicate", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterReplicate indicates an expected call of ClusterReplicate.
func (mr *MockCmdableMockRecorder) ClusterReplicate(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterReplicate", reflect.TypeOf((*MockCmdable)(nil).ClusterReplicate), ctx, nodeID)
}

// ClusterResetHard mocks base method.
func (m *MockCmdable) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetHard", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetHard indicates an expected call of ClusterResetHard.
func (mr *MockCmdableMockRecorder) ClusterResetHard(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetHard", reflect.TypeOf((*MockCmdable)(nil).ClusterResetHard), ctx)
}

// ClusterResetSoft mocks base method.
func (m *MockCmdable) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterResetSoft", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterResetSoft indicates an expected call of ClusterResetSoft.
func (mr *MockCmdableMockRecorder) ClusterResetSoft(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterResetSoft", reflect.TypeOf((*MockCmdable)(nil).ClusterResetSoft), ctx)
}

// ClusterSaveConfig mocks base method.
func (m *MockCmdable) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSaveConfig", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ClusterSaveConfig indicates an expected call of ClusterSaveConfig.
func (mr *MockCmdableMockRecorder) ClusterSaveConfig(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSaveConfig", reflect.TypeOf((*MockCmdable)(nil).ClusterSaveConfig), ctx)
}

// ClusterSlaves mocks base method.
func (m *MockCmdable) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlaves", ctx, nodeID)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ClusterSlaves indicates an expected call of ClusterSlaves.
func (mr *MockCmdableMockRecorder) ClusterSlaves(ctx, nodeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlaves", reflect.TypeOf((*MockCmdable)(nil).ClusterSlaves), ctx, nodeID)
}

// ClusterSlots mocks base method.
func (m *MockCmdable) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterSlots", ctx)
	ret0, _ := ret[0].(*redis.ClusterSlotsCmd)
	return ret0
}

// ClusterSlots indicates an expected call of ClusterSlots.
func (mr *MockCmdableMockRecorder) ClusterSlots(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterSlots", reflect.TypeOf((*MockCmdable)(nil).ClusterSlots), ctx)
}

// Command mocks base method.
func (m *MockCmdable) Command(ctx context.Context) *redis.CommandsInfoCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", ctx)
	ret0, _ := ret[0].(*redis.CommandsInfoCmd)
	return ret0
}

// Command indicates an expected call of Command.
func (mr *MockCmdableMockRecorder) Command(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCmdable)(nil).Command), ctx)
}

// ConfigGet mocks base method.
func (m *MockCmdable) ConfigGet(ctx context.Context, parameter string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigGet", ctx, parameter)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// ConfigGet indicates an expected call of ConfigGet.
func (mr *MockCmdableMockRecorder) ConfigGet(ctx, parameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigGet", reflect.TypeOf((*MockCmdable)(nil).ConfigGet), ctx, parameter)
}

// ConfigResetStat mocks base method.
func (m *MockCmdable) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigResetStat", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigResetStat indicates an expected call of ConfigResetStat.
func (mr *MockCmdableMockRecorder) ConfigResetStat(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigResetStat", reflect.TypeOf((*MockCmdable)(nil).ConfigResetStat), ctx)
}

// ConfigRewrite mocks base method.
func (m *MockCmdable) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigRewrite", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigRewrite indicates an expected call of ConfigRewrite.
func (mr *MockCmdableMockRecorder) ConfigRewrite(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigRewrite", reflect.TypeOf((*MockCmdable)(nil).ConfigRewrite), ctx)
}

// ConfigSet mocks base method.
func (m *MockCmdable) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigSet", ctx, parameter, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ConfigSet indicates an expected call of ConfigSet.
func (mr *MockCmdableMockRecorder) ConfigSet(ctx, parameter, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigSet", reflect.TypeOf((*MockCmdable)(nil).ConfigSet), ctx, parameter, value)
}

// Copy mocks base method.
func (m *MockCmdable) Copy(ctx context.Context, sourceKey, destKey string, db int, replace bool) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", ctx, sourceKey, destKey, db, replace)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockCmdableMockRecorder) Copy(ctx, sourceKey, destKey, db, replace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockCmdable)(nil).Copy), ctx, sourceKey, destKey, db, replace)
}

// DBSize mocks base method.
func (m *MockCmdable) DBSize(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBSize", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DBSize indicates an expected call of DBSize.
func (mr *MockCmdableMockRecorder) DBSize(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBSize", reflect.TypeOf((*MockCmdable)(nil).DBSize), ctx)
}

// DebugObject mocks base method.
func (m *MockCmdable) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DebugObject", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// DebugObject indicates an expected call of DebugObject.
func (mr *MockCmdableMockRecorder) DebugObject(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DebugObject", reflect.TypeOf((*MockCmdable)(nil).DebugObject), ctx, key)
}

// Decr mocks base method.
func (m *MockCmdable) Decr(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decr", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Decr indicates an expected call of Decr.
func (mr *MockCmdableMockRecorder) Decr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockCmdable)(nil).Decr), ctx, key)
}

// DecrBy mocks base method.
func (m *MockCmdable) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrBy", ctx, key, decrement)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// DecrBy indicates an expected call of DecrBy.
func (mr *MockCmdableMockRecorder) DecrBy(ctx, key, decrement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrBy", reflect.TypeOf((*MockCmdable)(nil).DecrBy), ctx, key, decrement)
}

// Del mocks base method.
func (m *MockCmdable) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockCmdableMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockCmdable)(nil).Del), varargs...)
}

// Dump mocks base method.
func (m *MockCmdable) Dump(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Dump indicates an expected call of Dump.
func (mr *MockCmdableMockRecorder) Dump(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockCmdable)(nil).Dump), ctx, key)
}

// Echo mocks base method.
func (m *MockCmdable) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Echo", ctx, message)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Echo indicates an expected call of Echo.
func (mr *MockCmdableMockRecorder) Echo(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Echo", reflect.TypeOf((*MockCmdable)(nil).Echo), ctx, message)
}

// Eval mocks base method.
func (m *MockCmdable) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, script, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eval", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// Eval indicates an expected call of Eval.
func (mr *MockCmdableMockRecorder) Eval(ctx, script, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, script, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockCmdable)(nil).Eval), varargs...)
}

// EvalSha mocks base method.
func (m *MockCmdable) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sha1, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalSha", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalSha indicates an expected call of EvalSha.
func (mr *MockCmdableMockRecorder) EvalSha(ctx, sha1, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sha1, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSha", reflect.TypeOf((*MockCmdable)(nil).EvalSha), varargs...)
}

// Exists mocks base method.
func (m *MockCmdable) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockCmdableMockRecorder) Exists(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockCmdable)(nil).Exists), varargs...)
}

// Expire mocks base method.
func (m *MockCmdable) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockCmdableMockRecorder) Expire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockCmdable)(nil).Expire), ctx, key, expiration)
}

// ExpireAt mocks base method.
func (m *MockCmdable) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireAt", ctx, key, tm)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireAt indicates an expected call of ExpireAt.
func (mr *MockCmdableMockRecorder) ExpireAt(ctx, key, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireAt", reflect.TypeOf((*MockCmdable)(nil).ExpireAt), ctx, key, tm)
}

// ExpireGT mocks base method.
func (m *MockCmdable) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireGT", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireGT indicates an expected call of ExpireGT.
func (mr *MockCmdableMockRecorder) ExpireGT(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireGT", reflect.TypeOf((*MockCmdable)(nil).ExpireGT), ctx, key, expiration)
}

// ExpireLT mocks base method.
func (m *MockCmdable) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireLT", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireLT indicates an expected call of ExpireLT.
func (mr *MockCmdableMockRecorder) ExpireLT(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireLT", reflect.TypeOf((*MockCmdable)(nil).ExpireLT), ctx, key, expiration)
}

// ExpireNX mocks base method.
func (m *MockCmdable) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireNX", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireNX indicates an expected call of ExpireNX.
func (mr *MockCmdableMockRecorder) ExpireNX(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireNX", reflect.TypeOf((*MockCmdable)(nil).ExpireNX), ctx, key, expiration)
}

// ExpireXX mocks base method.
func (m *MockCmdable) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpireXX", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// ExpireXX indicates an expected call of ExpireXX.
func (mr *MockCmdableMockRecorder) ExpireXX(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpireXX", reflect.TypeOf((*MockCmdable)(nil).ExpireXX), ctx, key, expiration)
}

// FlushAll mocks base method.
func (m *MockCmdable) FlushAll(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockCmdableMockRecorder) FlushAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockCmdable)(nil).FlushAll), ctx)
}

// FlushAllAsync mocks base method.
func (m *MockCmdable) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAllAsync", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushAllAsync indicates an expected call of FlushAllAsync.
func (mr *MockCmdableMockRecorder) FlushAllAsync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAllAsync", reflect.TypeOf((*MockCmdable)(nil).FlushAllAsync), ctx)
}

// FlushDB mocks base method.
func (m *MockCmdable) FlushDB(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDB", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDB indicates an expected call of FlushDB.
func (mr *MockCmdableMockRecorder) FlushDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDB", reflect.TypeOf((*MockCmdable)(nil).FlushDB), ctx)
}

// FlushDBAsync mocks base method.
func (m *MockCmdable) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushDBAsync", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// FlushDBAsync indicates an expected call of FlushDBAsync.
func (mr *MockCmdableMockRecorder) FlushDBAsync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushDBAsync", reflect.TypeOf((*MockCmdable)(nil).FlushDBAsync), ctx)
}

// GeoAdd mocks base method.
func (m *MockCmdable) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range geoLocation {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoAdd indicates an expected call of GeoAdd.
func (mr *MockCmdableMockRecorder) GeoAdd(ctx, key interface{}, geoLocation ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, geoLocation...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoAdd", reflect.TypeOf((*MockCmdable)(nil).GeoAdd), varargs...)
}

// GeoDist mocks base method.
func (m *MockCmdable) GeoDist(ctx context.Context, key, member1, member2, unit string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoDist", ctx, key, member1, member2, unit)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// GeoDist indicates an expected call of GeoDist.
func (mr *MockCmdableMockRecorder) GeoDist(ctx, key, member1, member2, unit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoDist", reflect.TypeOf((*MockCmdable)(nil).GeoDist), ctx, key, member1, member2, unit)
}

// GeoHash mocks base method.
func (m *MockCmdable) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoHash", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoHash indicates an expected call of GeoHash.
func (mr *MockCmdableMockRecorder) GeoHash(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoHash", reflect.TypeOf((*MockCmdable)(nil).GeoHash), varargs...)
}

// GeoPos mocks base method.
func (m *MockCmdable) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GeoPos", varargs...)
	ret0, _ := ret[0].(*redis.GeoPosCmd)
	return ret0
}

// GeoPos indicates an expected call of GeoPos.
func (mr *MockCmdableMockRecorder) GeoPos(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoPos", reflect.TypeOf((*MockCmdable)(nil).GeoPos), varargs...)
}

// GeoRadius mocks base method.
func (m *MockCmdable) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadius", ctx, key, longitude, latitude, query)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadius indicates an expected call of GeoRadius.
func (mr *MockCmdableMockRecorder) GeoRadius(ctx, key, longitude, latitude, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadius", reflect.TypeOf((*MockCmdable)(nil).GeoRadius), ctx, key, longitude, latitude, query)
}

// GeoRadiusByMember mocks base method.
func (m *MockCmdable) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMember", ctx, key, member, query)
	ret0, _ := ret[0].(*redis.GeoLocationCmd)
	return ret0
}

// GeoRadiusByMember indicates an expected call of GeoRadiusByMember.
func (mr *MockCmdableMockRecorder) GeoRadiusByMember(ctx, key, member, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMember", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusByMember), ctx, key, member, query)
}

// GeoRadiusByMemberStore mocks base method.
func (m *MockCmdable) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusByMemberStore", ctx, key, member, query)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusByMemberStore indicates an expected call of GeoRadiusByMemberStore.
func (mr *MockCmdableMockRecorder) GeoRadiusByMemberStore(ctx, key, member, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusByMemberStore", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusByMemberStore), ctx, key, member, query)
}

// GeoRadiusStore mocks base method.
func (m *MockCmdable) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoRadiusStore", ctx, key, longitude, latitude, query)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoRadiusStore indicates an expected call of GeoRadiusStore.
func (mr *MockCmdableMockRecorder) GeoRadiusStore(ctx, key, longitude, latitude, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoRadiusStore", reflect.TypeOf((*MockCmdable)(nil).GeoRadiusStore), ctx, key, longitude, latitude, query)
}

// GeoSearch mocks base method.
func (m *MockCmdable) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearch", ctx, key, q)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// GeoSearch indicates an expected call of GeoSearch.
func (mr *MockCmdableMockRecorder) GeoSearch(ctx, key, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearch", reflect.TypeOf((*MockCmdable)(nil).GeoSearch), ctx, key, q)
}

// GeoSearchLocation mocks base method.
func (m *MockCmdable) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchLocation", ctx, key, q)
	ret0, _ := ret[0].(*redis.GeoSearchLocationCmd)
	return ret0
}

// GeoSearchLocation indicates an expected call of GeoSearchLocation.
func (mr *MockCmdableMockRecorder) GeoSearchLocation(ctx, key, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchLocation", reflect.TypeOf((*MockCmdable)(nil).GeoSearchLocation), ctx, key, q)
}

// GeoSearchStore mocks base method.
func (m *MockCmdable) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeoSearchStore", ctx, key, store, q)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GeoSearchStore indicates an expected call of GeoSearchStore.
func (mr *MockCmdableMockRecorder) GeoSearchStore(ctx, key, store, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeoSearchStore", reflect.TypeOf((*MockCmdable)(nil).GeoSearchStore), ctx, key, store, q)
}

// Get mocks base method.
func (m *MockCmdable) Get(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCmdableMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCmdable)(nil).Get), ctx, key)
}

// GetBit mocks base method.
func (m *MockCmdable) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBit", ctx, key, offset)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// GetBit indicates an expected call of GetBit.
func (mr *MockCmdableMockRecorder) GetBit(ctx, key, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBit", reflect.TypeOf((*MockCmdable)(nil).GetBit), ctx, key, offset)
}

// GetDel mocks base method.
func (m *MockCmdable) GetDel(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDel", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetDel indicates an expected call of GetDel.
func (mr *MockCmdableMockRecorder) GetDel(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDel", reflect.TypeOf((*MockCmdable)(nil).GetDel), ctx, key)
}

// GetEx mocks base method.
func (m *MockCmdable) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEx", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetEx indicates an expected call of GetEx.
func (mr *MockCmdableMockRecorder) GetEx(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEx", reflect.TypeOf((*MockCmdable)(nil).GetEx), ctx, key, expiration)
}

// GetRange mocks base method.
func (m *MockCmdable) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRange", ctx, key, start, end)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetRange indicates an expected call of GetRange.
func (mr *MockCmdableMockRecorder) GetRange(ctx, key, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRange", reflect.TypeOf((*MockCmdable)(nil).GetRange), ctx, key, start, end)
}

// GetSet mocks base method.
func (m *MockCmdable) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSet", ctx, key, value)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// GetSet indicates an expected call of GetSet.
func (mr *MockCmdableMockRecorder) GetSet(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSet", reflect.TypeOf((*MockCmdable)(nil).GetSet), ctx, key, value)
}

// HDel mocks base method.
func (m *MockCmdable) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HDel indicates an expected call of HDel.
func (mr *MockCmdableMockRecorder) HDel(ctx, key interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockCmdable)(nil).HDel), varargs...)
}

// HExists mocks base method.
func (m *MockCmdable) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HExists", ctx, key, field)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HExists indicates an expected call of HExists.
func (mr *MockCmdableMockRecorder) HExists(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HExists", reflect.TypeOf((*MockCmdable)(nil).HExists), ctx, key, field)
}

// HGet mocks base method.
func (m *MockCmdable) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", ctx, key, field)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// HGet indicates an expected call of HGet.
func (mr *MockCmdableMockRecorder) HGet(ctx, key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockCmdable)(nil).HGet), ctx, key, field)
}

// HGetAll mocks base method.
func (m *MockCmdable) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", ctx, key)
	ret0, _ := ret[0].(*redis.StringStringMapCmd)
	return ret0
}

// HGetAll indicates an expected call of HGetAll.
func (mr *MockCmdableMockRecorder) HGetAll(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockCmdable)(nil).HGetAll), ctx, key)
}

// HIncrBy mocks base method.
func (m *MockCmdable) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrBy", ctx, key, field, incr)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HIncrBy indicates an expected call of HIncrBy.
func (mr *MockCmdableMockRecorder) HIncrBy(ctx, key, field, incr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrBy", reflect.TypeOf((*MockCmdable)(nil).HIncrBy), ctx, key, field, incr)
}

// HIncrByFloat mocks base method.
func (m *MockCmdable) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HIncrByFloat", ctx, key, field, incr)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// HIncrByFloat indicates an expected call of HIncrByFloat.
func (mr *MockCmdableMockRecorder) HIncrByFloat(ctx, key, field, incr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HIncrByFloat", reflect.TypeOf((*MockCmdable)(nil).HIncrByFloat), ctx, key, field, incr)
}

// HKeys mocks base method.
func (m *MockCmdable) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HKeys", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HKeys indicates an expected call of HKeys.
func (mr *MockCmdableMockRecorder) HKeys(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HKeys", reflect.TypeOf((*MockCmdable)(nil).HKeys), ctx, key)
}

// HLen mocks base method.
func (m *MockCmdable) HLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HLen indicates an expected call of HLen.
func (mr *MockCmdableMockRecorder) HLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HLen", reflect.TypeOf((*MockCmdable)(nil).HLen), ctx, key)
}

// HMGet mocks base method.
func (m *MockCmdable) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// HMGet indicates an expected call of HMGet.
func (mr *MockCmdableMockRecorder) HMGet(ctx, key interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockCmdable)(nil).HMGet), varargs...)
}

// HMSet mocks base method.
func (m *MockCmdable) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMSet", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HMSet indicates an expected call of HMSet.
func (mr *MockCmdableMockRecorder) HMSet(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockCmdable)(nil).HMSet), varargs...)
}

// HRandField mocks base method.
func (m *MockCmdable) HRandField(ctx context.Context, key string, count int, withValues bool) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HRandField", ctx, key, count, withValues)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HRandField indicates an expected call of HRandField.
func (mr *MockCmdableMockRecorder) HRandField(ctx, key, count, withValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HRandField", reflect.TypeOf((*MockCmdable)(nil).HRandField), ctx, key, count, withValues)
}

// HScan mocks base method.
func (m *MockCmdable) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// HScan indicates an expected call of HScan.
func (mr *MockCmdableMockRecorder) HScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HScan", reflect.TypeOf((*MockCmdable)(nil).HScan), ctx, key, cursor, match, count)
}

// HSet mocks base method.
func (m *MockCmdable) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HSet", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockCmdableMockRecorder) HSet(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockCmdable)(nil).HSet), varargs...)
}

// HSetNX mocks base method.
func (m *MockCmdable) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSetNX", ctx, key, field, value)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HSetNX indicates an expected call of HSetNX.
func (mr *MockCmdableMockRecorder) HSetNX(ctx, key, field, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSetNX", reflect.TypeOf((*MockCmdable)(nil).HSetNX), ctx, key, field, value)
}

// HVals mocks base method.
func (m *MockCmdable) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HVals", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// HVals indicates an expected call of HVals.
func (mr *MockCmdableMockRecorder) HVals(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HVals", reflect.TypeOf((*MockCmdable)(nil).HVals), ctx, key)
}

// Incr mocks base method.
func (m *MockCmdable) Incr(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Incr", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockCmdableMockRecorder) Incr(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockCmdable)(nil).Incr), ctx, key)
}

// IncrBy mocks base method.
func (m *MockCmdable) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrBy", ctx, key, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// IncrBy indicates an expected call of IncrBy.
func (mr *MockCmdableMockRecorder) IncrBy(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrBy", reflect.TypeOf((*MockCmdable)(nil).IncrBy), ctx, key, value)
}

// IncrByFloat mocks base method.
func (m *MockCmdable) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrByFloat", ctx, key, value)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// IncrByFloat indicates an expected call of IncrByFloat.
func (mr *MockCmdableMockRecorder) IncrByFloat(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrByFloat", reflect.TypeOf((*MockCmdable)(nil).IncrByFloat), ctx, key, value)
}

// Info mocks base method.
func (m *MockCmdable) Info(ctx context.Context, section ...string) *redis.StringCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range section {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockCmdableMockRecorder) Info(ctx interface{}, section ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, section...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockCmdable)(nil).Info), varargs...)
}

// Keys mocks base method.
func (m *MockCmdable) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", ctx, pattern)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockCmdableMockRecorder) Keys(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockCmdable)(nil).Keys), ctx, pattern)
}

// LIndex mocks base method.
func (m *MockCmdable) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LIndex", ctx, key, index)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LIndex indicates an expected call of LIndex.
func (mr *MockCmdableMockRecorder) LIndex(ctx, key, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LIndex", reflect.TypeOf((*MockCmdable)(nil).LIndex), ctx, key, index)
}

// LInsert mocks base method.
func (m *MockCmdable) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsert", ctx, key, op, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsert indicates an expected call of LInsert.
func (mr *MockCmdableMockRecorder) LInsert(ctx, key, op, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsert", reflect.TypeOf((*MockCmdable)(nil).LInsert), ctx, key, op, pivot, value)
}

// LInsertAfter mocks base method.
func (m *MockCmdable) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertAfter", ctx, key, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertAfter indicates an expected call of LInsertAfter.
func (mr *MockCmdableMockRecorder) LInsertAfter(ctx, key, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertAfter", reflect.TypeOf((*MockCmdable)(nil).LInsertAfter), ctx, key, pivot, value)
}

// LInsertBefore mocks base method.
func (m *MockCmdable) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LInsertBefore", ctx, key, pivot, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LInsertBefore indicates an expected call of LInsertBefore.
func (mr *MockCmdableMockRecorder) LInsertBefore(ctx, key, pivot, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LInsertBefore", reflect.TypeOf((*MockCmdable)(nil).LInsertBefore), ctx, key, pivot, value)
}

// LLen mocks base method.
func (m *MockCmdable) LLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LLen indicates an expected call of LLen.
func (mr *MockCmdableMockRecorder) LLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LLen", reflect.TypeOf((*MockCmdable)(nil).LLen), ctx, key)
}

// LMove mocks base method.
func (m *MockCmdable) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LMove", ctx, source, destination, srcpos, destpos)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LMove indicates an expected call of LMove.
func (mr *MockCmdableMockRecorder) LMove(ctx, source, destination, srcpos, destpos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LMove", reflect.TypeOf((*MockCmdable)(nil).LMove), ctx, source, destination, srcpos, destpos)
}

// LPop mocks base method.
func (m *MockCmdable) LPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// LPop indicates an expected call of LPop.
func (mr *MockCmdableMockRecorder) LPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPop", reflect.TypeOf((*MockCmdable)(nil).LPop), ctx, key)
}

// LPopCount mocks base method.
func (m *MockCmdable) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPopCount", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LPopCount indicates an expected call of LPopCount.
func (mr *MockCmdableMockRecorder) LPopCount(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPopCount", reflect.TypeOf((*MockCmdable)(nil).LPopCount), ctx, key, count)
}

// LPos mocks base method.
func (m *MockCmdable) LPos(ctx context.Context, key, value string, args redis.LPosArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPos", ctx, key, value, args)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPos indicates an expected call of LPos.
func (mr *MockCmdableMockRecorder) LPos(ctx, key, value, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPos", reflect.TypeOf((*MockCmdable)(nil).LPos), ctx, key, value, args)
}

// LPosCount mocks base method.
func (m *MockCmdable) LPosCount(ctx context.Context, key, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LPosCount", ctx, key, value, count, args)
	ret0, _ := ret[0].(*redis.IntSliceCmd)
	return ret0
}

// LPosCount indicates an expected call of LPosCount.
func (mr *MockCmdableMockRecorder) LPosCount(ctx, key, value, count, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPosCount", reflect.TypeOf((*MockCmdable)(nil).LPosCount), ctx, key, value, count, args)
}

// LPush mocks base method.
func (m *MockCmdable) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPush indicates an expected call of LPush.
func (mr *MockCmdableMockRecorder) LPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPush", reflect.TypeOf((*MockCmdable)(nil).LPush), varargs...)
}

// LPushX mocks base method.
func (m *MockCmdable) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LPushX indicates an expected call of LPushX.
func (mr *MockCmdableMockRecorder) LPushX(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LPushX", reflect.TypeOf((*MockCmdable)(nil).LPushX), varargs...)
}

// LRange mocks base method.
func (m *MockCmdable) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LRange indicates an expected call of LRange.
func (mr *MockCmdableMockRecorder) LRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRange", reflect.TypeOf((*MockCmdable)(nil).LRange), ctx, key, start, stop)
}

// LRem mocks base method.
func (m *MockCmdable) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRem", ctx, key, count, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LRem indicates an expected call of LRem.
func (mr *MockCmdableMockRecorder) LRem(ctx, key, count, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRem", reflect.TypeOf((*MockCmdable)(nil).LRem), ctx, key, count, value)
}

// LSet mocks base method.
func (m *MockCmdable) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LSet", ctx, key, index, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LSet indicates an expected call of LSet.
func (mr *MockCmdableMockRecorder) LSet(ctx, key, index, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LSet", reflect.TypeOf((*MockCmdable)(nil).LSet), ctx, key, index, value)
}

// LTrim mocks base method.
func (m *MockCmdable) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LTrim", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// LTrim indicates an expected call of LTrim.
func (mr *MockCmdableMockRecorder) LTrim(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LTrim", reflect.TypeOf((*MockCmdable)(nil).LTrim), ctx, key, start, stop)
}

// LastSave mocks base method.
func (m *MockCmdable) LastSave(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSave", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// LastSave indicates an expected call of LastSave.
func (mr *MockCmdableMockRecorder) LastSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSave", reflect.TypeOf((*MockCmdable)(nil).LastSave), ctx)
}

// MGet mocks base method.
func (m *MockCmdable) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// MGet indicates an expected call of MGet.
func (mr *MockCmdableMockRecorder) MGet(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockCmdable)(nil).MGet), varargs...)
}

// MSet mocks base method.
func (m *MockCmdable) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSet", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// MSet indicates an expected call of MSet.
func (mr *MockCmdableMockRecorder) MSet(ctx interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSet", reflect.TypeOf((*MockCmdable)(nil).MSet), varargs...)
}

// MSetNX mocks base method.
func (m *MockCmdable) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MSetNX", varargs...)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// MSetNX indicates an expected call of MSetNX.
func (mr *MockCmdableMockRecorder) MSetNX(ctx interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MSetNX", reflect.TypeOf((*MockCmdable)(nil).MSetNX), varargs...)
}

// MemoryUsage mocks base method.
func (m *MockCmdable) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range samples {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MemoryUsage", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// MemoryUsage indicates an expected call of MemoryUsage.
func (mr *MockCmdableMockRecorder) MemoryUsage(ctx, key interface{}, samples ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, samples...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MemoryUsage", reflect.TypeOf((*MockCmdable)(nil).MemoryUsage), varargs...)
}

// Migrate mocks base method.
func (m *MockCmdable) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrate", ctx, host, port, key, db, timeout)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockCmdableMockRecorder) Migrate(ctx, host, port, key, db, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockCmdable)(nil).Migrate), ctx, host, port, key, db, timeout)
}

// Move mocks base method.
func (m *MockCmdable) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", ctx, key, db)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockCmdableMockRecorder) Move(ctx, key, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockCmdable)(nil).Move), ctx, key, db)
}

// ObjectEncoding mocks base method.
func (m *MockCmdable) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectEncoding", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ObjectEncoding indicates an expected call of ObjectEncoding.
func (mr *MockCmdableMockRecorder) ObjectEncoding(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectEncoding", reflect.TypeOf((*MockCmdable)(nil).ObjectEncoding), ctx, key)
}

// ObjectIdleTime mocks base method.
func (m *MockCmdable) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectIdleTime", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// ObjectIdleTime indicates an expected call of ObjectIdleTime.
func (mr *MockCmdableMockRecorder) ObjectIdleTime(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectIdleTime", reflect.TypeOf((*MockCmdable)(nil).ObjectIdleTime), ctx, key)
}

// ObjectRefCount mocks base method.
func (m *MockCmdable) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectRefCount", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ObjectRefCount indicates an expected call of ObjectRefCount.
func (mr *MockCmdableMockRecorder) ObjectRefCount(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectRefCount", reflect.TypeOf((*MockCmdable)(nil).ObjectRefCount), ctx, key)
}

// PExpire mocks base method.
func (m *MockCmdable) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpire", ctx, key, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpire indicates an expected call of PExpire.
func (mr *MockCmdableMockRecorder) PExpire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpire", reflect.TypeOf((*MockCmdable)(nil).PExpire), ctx, key, expiration)
}

// PExpireAt mocks base method.
func (m *MockCmdable) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PExpireAt", ctx, key, tm)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// PExpireAt indicates an expected call of PExpireAt.
func (mr *MockCmdableMockRecorder) PExpireAt(ctx, key, tm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PExpireAt", reflect.TypeOf((*MockCmdable)(nil).PExpireAt), ctx, key, tm)
}

// PFAdd mocks base method.
func (m *MockCmdable) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range els {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFAdd indicates an expected call of PFAdd.
func (mr *MockCmdableMockRecorder) PFAdd(ctx, key interface{}, els ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, els...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFAdd", reflect.TypeOf((*MockCmdable)(nil).PFAdd), varargs...)
}

// PFCount mocks base method.
func (m *MockCmdable) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFCount", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PFCount indicates an expected call of PFCount.
func (mr *MockCmdableMockRecorder) PFCount(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFCount", reflect.TypeOf((*MockCmdable)(nil).PFCount), varargs...)
}

// PFMerge mocks base method.
func (m *MockCmdable) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PFMerge", varargs...)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// PFMerge indicates an expected call of PFMerge.
func (mr *MockCmdableMockRecorder) PFMerge(ctx, dest interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PFMerge", reflect.TypeOf((*MockCmdable)(nil).PFMerge), varargs...)
}

// PTTL mocks base method.
func (m *MockCmdable) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PTTL", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// PTTL indicates an expected call of PTTL.
func (mr *MockCmdableMockRecorder) PTTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PTTL", reflect.TypeOf((*MockCmdable)(nil).PTTL), ctx, key)
}

// Persist mocks base method.
func (m *MockCmdable) Persist(ctx context.Context, key string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", ctx, key)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockCmdableMockRecorder) Persist(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockCmdable)(nil).Persist), ctx, key)
}

// Ping mocks base method.
func (m *MockCmdable) Ping(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockCmdableMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCmdable)(nil).Ping), ctx)
}

// Pipeline mocks base method.
func (m *MockCmdable) Pipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockCmdableMockRecorder) Pipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockCmdable)(nil).Pipeline))
}

// Pipelined mocks base method.
func (m *MockCmdable) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelined", ctx, fn)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelined indicates an expected call of Pipelined.
func (mr *MockCmdableMockRecorder) Pipelined(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelined", reflect.TypeOf((*MockCmdable)(nil).Pipelined), ctx, fn)
}

// PubSubChannels mocks base method.
func (m *MockCmdable) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubChannels", ctx, pattern)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// PubSubChannels indicates an expected call of PubSubChannels.
func (mr *MockCmdableMockRecorder) PubSubChannels(ctx, pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubChannels", reflect.TypeOf((*MockCmdable)(nil).PubSubChannels), ctx, pattern)
}

// PubSubNumPat mocks base method.
func (m *MockCmdable) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubSubNumPat", ctx)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// PubSubNumPat indicates an expected call of PubSubNumPat.
func (mr *MockCmdableMockRecorder) PubSubNumPat(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumPat", reflect.TypeOf((*MockCmdable)(nil).PubSubNumPat), ctx)
}

// PubSubNumSub mocks base method.
func (m *MockCmdable) PubSubNumSub(ctx context.Context, channels ...string) *redis.StringIntMapCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range channels {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PubSubNumSub", varargs...)
	ret0, _ := ret[0].(*redis.StringIntMapCmd)
	return ret0
}

// PubSubNumSub indicates an expected call of PubSubNumSub.
func (mr *MockCmdableMockRecorder) PubSubNumSub(ctx interface{}, channels ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, channels...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubSubNumSub", reflect.TypeOf((*MockCmdable)(nil).PubSubNumSub), varargs...)
}

// Publish mocks base method.
func (m *MockCmdable) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, channel, message)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockCmdableMockRecorder) Publish(ctx, channel, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockCmdable)(nil).Publish), ctx, channel, message)
}

// Quit mocks base method.
func (m *MockCmdable) Quit(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockCmdableMockRecorder) Quit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockCmdable)(nil).Quit), ctx)
}

// RPop mocks base method.
func (m *MockCmdable) RPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPop indicates an expected call of RPop.
func (mr *MockCmdableMockRecorder) RPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPop", reflect.TypeOf((*MockCmdable)(nil).RPop), ctx, key)
}

// RPopCount mocks base method.
func (m *MockCmdable) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopCount", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// RPopCount indicates an expected call of RPopCount.
func (mr *MockCmdableMockRecorder) RPopCount(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopCount", reflect.TypeOf((*MockCmdable)(nil).RPopCount), ctx, key, count)
}

// RPopLPush mocks base method.
func (m *MockCmdable) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopLPush", ctx, source, destination)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPopLPush indicates an expected call of RPopLPush.
func (mr *MockCmdableMockRecorder) RPopLPush(ctx, source, destination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopLPush", reflect.TypeOf((*MockCmdable)(nil).RPopLPush), ctx, source, destination)
}

// RPush mocks base method.
func (m *MockCmdable) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPush indicates an expected call of RPush.
func (mr *MockCmdableMockRecorder) RPush(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockCmdable)(nil).RPush), varargs...)
}

// RPushX mocks base method.
func (m *MockCmdable) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPushX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPushX indicates an expected call of RPushX.
func (mr *MockCmdableMockRecorder) RPushX(ctx, key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPushX", reflect.TypeOf((*MockCmdable)(nil).RPushX), varargs...)
}

// RandomKey mocks base method.
func (m *MockCmdable) RandomKey(ctx context.Context) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RandomKey", ctx)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RandomKey indicates an expected call of RandomKey.
func (mr *MockCmdableMockRecorder) RandomKey(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RandomKey", reflect.TypeOf((*MockCmdable)(nil).RandomKey), ctx)
}

// ReadOnly mocks base method.
func (m *MockCmdable) ReadOnly(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOnly", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadOnly indicates an expected call of ReadOnly.
func (mr *MockCmdableMockRecorder) ReadOnly(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOnly", reflect.TypeOf((*MockCmdable)(nil).ReadOnly), ctx)
}

// ReadWrite mocks base method.
func (m *MockCmdable) ReadWrite(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWrite", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ReadWrite indicates an expected call of ReadWrite.
func (mr *MockCmdableMockRecorder) ReadWrite(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWrite", reflect.TypeOf((*MockCmdable)(nil).ReadWrite), ctx)
}

// Rename mocks base method.
func (m *MockCmdable) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", ctx, key, newkey)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockCmdableMockRecorder) Rename(ctx, key, newkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockCmdable)(nil).Rename), ctx, key, newkey)
}

// RenameNX mocks base method.
func (m *MockCmdable) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameNX", ctx, key, newkey)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// RenameNX indicates an expected call of RenameNX.
func (mr *MockCmdableMockRecorder) RenameNX(ctx, key, newkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameNX", reflect.TypeOf((*MockCmdable)(nil).RenameNX), ctx, key, newkey)
}

// Restore mocks base method.
func (m *MockCmdable) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, key, ttl, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockCmdableMockRecorder) Restore(ctx, key, ttl, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockCmdable)(nil).Restore), ctx, key, ttl, value)
}

// RestoreReplace mocks base method.
func (m *MockCmdable) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreReplace", ctx, key, ttl, value)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// RestoreReplace indicates an expected call of RestoreReplace.
func (mr *MockCmdableMockRecorder) RestoreReplace(ctx, key, ttl, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreReplace", reflect.TypeOf((*MockCmdable)(nil).RestoreReplace), ctx, key, ttl, value)
}

// SAdd mocks base method.
func (m *MockCmdable) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockCmdableMockRecorder) SAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockCmdable)(nil).SAdd), varargs...)
}

// SCard mocks base method.
func (m *MockCmdable) SCard(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SCard", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SCard indicates an expected call of SCard.
func (mr *MockCmdableMockRecorder) SCard(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SCard", reflect.TypeOf((*MockCmdable)(nil).SCard), ctx, key)
}

// SDiff mocks base method.
func (m *MockCmdable) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SDiff indicates an expected call of SDiff.
func (mr *MockCmdableMockRecorder) SDiff(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiff", reflect.TypeOf((*MockCmdable)(nil).SDiff), varargs...)
}

// SDiffStore mocks base method.
func (m *MockCmdable) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SDiffStore indicates an expected call of SDiffStore.
func (mr *MockCmdableMockRecorder) SDiffStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SDiffStore", reflect.TypeOf((*MockCmdable)(nil).SDiffStore), varargs...)
}

// SInter mocks base method.
func (m *MockCmdable) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInter", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SInter indicates an expected call of SInter.
func (mr *MockCmdableMockRecorder) SInter(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInter", reflect.TypeOf((*MockCmdable)(nil).SInter), varargs...)
}

// SInterStore mocks base method.
func (m *MockCmdable) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SInterStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SInterStore indicates an expected call of SInterStore.
func (mr *MockCmdableMockRecorder) SInterStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SInterStore", reflect.TypeOf((*MockCmdable)(nil).SInterStore), varargs...)
}

// SIsMember mocks base method.
func (m *MockCmdable) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIsMember", ctx, key, member)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SIsMember indicates an expected call of SIsMember.
func (mr *MockCmdableMockRecorder) SIsMember(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIsMember", reflect.TypeOf((*MockCmdable)(nil).SIsMember), ctx, key, member)
}

// SMIsMember mocks base method.
func (m *MockCmdable) SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SMIsMember", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// SMIsMember indicates an expected call of SMIsMember.
func (mr *MockCmdableMockRecorder) SMIsMember(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMIsMember", reflect.TypeOf((*MockCmdable)(nil).SMIsMember), varargs...)
}

// SMembers mocks base method.
func (m *MockCmdable) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", ctx, key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SMembers indicates an expected call of SMembers.
func (mr *MockCmdableMockRecorder) SMembers(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockCmdable)(nil).SMembers), ctx, key)
}

// SMembersMap mocks base method.
func (m *MockCmdable) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembersMap", ctx, key)
	ret0, _ := ret[0].(*redis.StringStructMapCmd)
	return ret0
}

// SMembersMap indicates an expected call of SMembersMap.
func (mr *MockCmdableMockRecorder) SMembersMap(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembersMap", reflect.TypeOf((*MockCmdable)(nil).SMembersMap), ctx, key)
}

// SMove mocks base method.
func (m *MockCmdable) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMove", ctx, source, destination, member)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SMove indicates an expected call of SMove.
func (mr *MockCmdableMockRecorder) SMove(ctx, source, destination, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMove", reflect.TypeOf((*MockCmdable)(nil).SMove), ctx, source, destination, member)
}

// SPop mocks base method.
func (m *MockCmdable) SPop(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPop", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SPop indicates an expected call of SPop.
func (mr *MockCmdableMockRecorder) SPop(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPop", reflect.TypeOf((*MockCmdable)(nil).SPop), ctx, key)
}

// SPopN mocks base method.
func (m *MockCmdable) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPopN", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SPopN indicates an expected call of SPopN.
func (mr *MockCmdableMockRecorder) SPopN(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPopN", reflect.TypeOf((*MockCmdable)(nil).SPopN), ctx, key, count)
}

// SRandMember mocks base method.
func (m *MockCmdable) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMember", ctx, key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// SRandMember indicates an expected call of SRandMember.
func (mr *MockCmdableMockRecorder) SRandMember(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMember", reflect.TypeOf((*MockCmdable)(nil).SRandMember), ctx, key)
}

// SRandMemberN mocks base method.
func (m *MockCmdable) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SRandMemberN", ctx, key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SRandMemberN indicates an expected call of SRandMemberN.
func (mr *MockCmdableMockRecorder) SRandMemberN(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRandMemberN", reflect.TypeOf((*MockCmdable)(nil).SRandMemberN), ctx, key, count)
}

// SRem mocks base method.
func (m *MockCmdable) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SRem indicates an expected call of SRem.
func (mr *MockCmdableMockRecorder) SRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockCmdable)(nil).SRem), varargs...)
}

// SScan mocks base method.
func (m *MockCmdable) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// SScan indicates an expected call of SScan.
func (mr *MockCmdableMockRecorder) SScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SScan", reflect.TypeOf((*MockCmdable)(nil).SScan), ctx, key, cursor, match, count)
}

// SUnion mocks base method.
func (m *MockCmdable) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnion", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SUnion indicates an expected call of SUnion.
func (mr *MockCmdableMockRecorder) SUnion(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnion", reflect.TypeOf((*MockCmdable)(nil).SUnion), varargs...)
}

// SUnionStore mocks base method.
func (m *MockCmdable) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SUnionStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SUnionStore indicates an expected call of SUnionStore.
func (mr *MockCmdableMockRecorder) SUnionStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SUnionStore", reflect.TypeOf((*MockCmdable)(nil).SUnionStore), varargs...)
}

// Save mocks base method.
func (m *MockCmdable) Save(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockCmdableMockRecorder) Save(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCmdable)(nil).Save), ctx)
}

// Scan mocks base method.
func (m *MockCmdable) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", ctx, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockCmdableMockRecorder) Scan(ctx, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockCmdable)(nil).Scan), ctx, cursor, match, count)
}

// ScanType mocks base method.
func (m *MockCmdable) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanType", ctx, cursor, match, count, keyType)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ScanType indicates an expected call of ScanType.
func (mr *MockCmdableMockRecorder) ScanType(ctx, cursor, match, count, keyType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanType", reflect.TypeOf((*MockCmdable)(nil).ScanType), ctx, cursor, match, count, keyType)
}

// ScriptExists mocks base method.
func (m *MockCmdable) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range hashes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScriptExists", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// ScriptExists indicates an expected call of ScriptExists.
func (mr *MockCmdableMockRecorder) ScriptExists(ctx interface{}, hashes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, hashes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptExists", reflect.TypeOf((*MockCmdable)(nil).ScriptExists), varargs...)
}

// ScriptFlush mocks base method.
func (m *MockCmdable) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptFlush", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptFlush indicates an expected call of ScriptFlush.
func (mr *MockCmdableMockRecorder) ScriptFlush(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptFlush", reflect.TypeOf((*MockCmdable)(nil).ScriptFlush), ctx)
}

// ScriptKill mocks base method.
func (m *MockCmdable) ScriptKill(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptKill", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ScriptKill indicates an expected call of ScriptKill.
func (mr *MockCmdableMockRecorder) ScriptKill(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptKill", reflect.TypeOf((*MockCmdable)(nil).ScriptKill), ctx)
}

// ScriptLoad mocks base method.
func (m *MockCmdable) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptLoad", ctx, script)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ScriptLoad indicates an expected call of ScriptLoad.
func (mr *MockCmdableMockRecorder) ScriptLoad(ctx, script interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptLoad", reflect.TypeOf((*MockCmdable)(nil).ScriptLoad), ctx, script)
}

// Set mocks base method.
func (m *MockCmdable) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCmdableMockRecorder) Set(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCmdable)(nil).Set), ctx, key, value, expiration)
}

// SetArgs mocks base method.
func (m *MockCmdable) SetArgs(ctx context.Context, key string, value interface{}, a redis.SetArgs) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetArgs", ctx, key, value, a)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetArgs indicates an expected call of SetArgs.
func (mr *MockCmdableMockRecorder) SetArgs(ctx, key, value, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetArgs", reflect.TypeOf((*MockCmdable)(nil).SetArgs), ctx, key, value, a)
}

// SetBit mocks base method.
func (m *MockCmdable) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBit", ctx, key, offset, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetBit indicates an expected call of SetBit.
func (mr *MockCmdableMockRecorder) SetBit(ctx, key, offset, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBit", reflect.TypeOf((*MockCmdable)(nil).SetBit), ctx, key, offset, value)
}

// SetEX mocks base method.
func (m *MockCmdable) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SetEX indicates an expected call of SetEX.
func (mr *MockCmdableMockRecorder) SetEX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEX", reflect.TypeOf((*MockCmdable)(nil).SetEX), ctx, key, value, expiration)
}

// SetNX mocks base method.
func (m *MockCmdable) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetNX indicates an expected call of SetNX.
func (mr *MockCmdableMockRecorder) SetNX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockCmdable)(nil).SetNX), ctx, key, value, expiration)
}

// SetRange mocks base method.
func (m *MockCmdable) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRange", ctx, key, offset, value)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SetRange indicates an expected call of SetRange.
func (mr *MockCmdableMockRecorder) SetRange(ctx, key, offset, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRange", reflect.TypeOf((*MockCmdable)(nil).SetRange), ctx, key, offset, value)
}

// SetXX mocks base method.
func (m *MockCmdable) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetXX", ctx, key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetXX indicates an expected call of SetXX.
func (mr *MockCmdableMockRecorder) SetXX(ctx, key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetXX", reflect.TypeOf((*MockCmdable)(nil).SetXX), ctx, key, value, expiration)
}

// Shutdown mocks base method.
func (m *MockCmdable) Shutdown(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockCmdableMockRecorder) Shutdown(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCmdable)(nil).Shutdown), ctx)
}

// ShutdownNoSave mocks base method.
func (m *MockCmdable) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownNoSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownNoSave indicates an expected call of ShutdownNoSave.
func (mr *MockCmdableMockRecorder) ShutdownNoSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownNoSave", reflect.TypeOf((*MockCmdable)(nil).ShutdownNoSave), ctx)
}

// ShutdownSave mocks base method.
func (m *MockCmdable) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownSave", ctx)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// ShutdownSave indicates an expected call of ShutdownSave.
func (mr *MockCmdableMockRecorder) ShutdownSave(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownSave", reflect.TypeOf((*MockCmdable)(nil).ShutdownSave), ctx)
}

// SlaveOf mocks base method.
func (m *MockCmdable) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlaveOf", ctx, host, port)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// SlaveOf indicates an expected call of SlaveOf.
func (mr *MockCmdableMockRecorder) SlaveOf(ctx, host, port interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlaveOf", reflect.TypeOf((*MockCmdable)(nil).SlaveOf), ctx, host, port)
}

// Sort mocks base method.
func (m *MockCmdable) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sort", ctx, key, sort)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// Sort indicates an expected call of Sort.
func (mr *MockCmdableMockRecorder) Sort(ctx, key, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockCmdable)(nil).Sort), ctx, key, sort)
}

// SortInterfaces mocks base method.
func (m *MockCmdable) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortInterfaces", ctx, key, sort)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// SortInterfaces indicates an expected call of SortInterfaces.
func (mr *MockCmdableMockRecorder) SortInterfaces(ctx, key, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortInterfaces", reflect.TypeOf((*MockCmdable)(nil).SortInterfaces), ctx, key, sort)
}

// SortStore mocks base method.
func (m *MockCmdable) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SortStore", ctx, key, store, sort)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SortStore indicates an expected call of SortStore.
func (mr *MockCmdableMockRecorder) SortStore(ctx, key, store, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortStore", reflect.TypeOf((*MockCmdable)(nil).SortStore), ctx, key, store, sort)
}

// StrLen mocks base method.
func (m *MockCmdable) StrLen(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StrLen", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// StrLen indicates an expected call of StrLen.
func (mr *MockCmdableMockRecorder) StrLen(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StrLen", reflect.TypeOf((*MockCmdable)(nil).StrLen), ctx, key)
}

// TTL mocks base method.
func (m *MockCmdable) TTL(ctx context.Context, key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", ctx, key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockCmdableMockRecorder) TTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockCmdable)(nil).TTL), ctx, key)
}

// Time mocks base method.
func (m *MockCmdable) Time(ctx context.Context) *redis.TimeCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time", ctx)
	ret0, _ := ret[0].(*redis.TimeCmd)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockCmdableMockRecorder) Time(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockCmdable)(nil).Time), ctx)
}

// Touch mocks base method.
func (m *MockCmdable) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Touch", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Touch indicates an expected call of Touch.
func (mr *MockCmdableMockRecorder) Touch(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Touch", reflect.TypeOf((*MockCmdable)(nil).Touch), varargs...)
}

// TxPipeline mocks base method.
func (m *MockCmdable) TxPipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// TxPipeline indicates an expected call of TxPipeline.
func (mr *MockCmdableMockRecorder) TxPipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipeline", reflect.TypeOf((*MockCmdable)(nil).TxPipeline))
}

// TxPipelined mocks base method.
func (m *MockCmdable) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipelined", ctx, fn)
	ret0, _ := ret[0].([]redis.Cmder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxPipelined indicates an expected call of TxPipelined.
func (mr *MockCmdableMockRecorder) TxPipelined(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipelined", reflect.TypeOf((*MockCmdable)(nil).TxPipelined), ctx, fn)
}

// Type mocks base method.
func (m *MockCmdable) Type(ctx context.Context, key string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type", ctx, key)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockCmdableMockRecorder) Type(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockCmdable)(nil).Type), ctx, key)
}

// Unlink mocks base method.
func (m *MockCmdable) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unlink", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Unlink indicates an expected call of Unlink.
func (mr *MockCmdableMockRecorder) Unlink(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlink", reflect.TypeOf((*MockCmdable)(nil).Unlink), varargs...)
}

// XAck mocks base method.
func (m *MockCmdable) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, stream, group}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XAck", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XAck indicates an expected call of XAck.
func (mr *MockCmdableMockRecorder) XAck(ctx, stream, group interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, stream, group}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAck", reflect.TypeOf((*MockCmdable)(nil).XAck), varargs...)
}

// XAdd mocks base method.
func (m *MockCmdable) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAdd", ctx, a)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// XAdd indicates an expected call of XAdd.
func (mr *MockCmdableMockRecorder) XAdd(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAdd", reflect.TypeOf((*MockCmdable)(nil).XAdd), ctx, a)
}

// XAutoClaim mocks base method.
func (m *MockCmdable) XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaim", ctx, a)
	ret0, _ := ret[0].(*redis.XAutoClaimCmd)
	return ret0
}

// XAutoClaim indicates an expected call of XAutoClaim.
func (mr *MockCmdableMockRecorder) XAutoClaim(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaim", reflect.TypeOf((*MockCmdable)(nil).XAutoClaim), ctx, a)
}

// XAutoClaimJustID mocks base method.
func (m *MockCmdable) XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XAutoClaimJustID", ctx, a)
	ret0, _ := ret[0].(*redis.XAutoClaimJustIDCmd)
	return ret0
}

// XAutoClaimJustID indicates an expected call of XAutoClaimJustID.
func (mr *MockCmdableMockRecorder) XAutoClaimJustID(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XAutoClaimJustID", reflect.TypeOf((*MockCmdable)(nil).XAutoClaimJustID), ctx, a)
}

// XClaim mocks base method.
func (m *MockCmdable) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaim", ctx, a)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XClaim indicates an expected call of XClaim.
func (mr *MockCmdableMockRecorder) XClaim(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaim", reflect.TypeOf((*MockCmdable)(nil).XClaim), ctx, a)
}

// XClaimJustID mocks base method.
func (m *MockCmdable) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XClaimJustID", ctx, a)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// XClaimJustID indicates an expected call of XClaimJustID.
func (mr *MockCmdableMockRecorder) XClaimJustID(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XClaimJustID", reflect.TypeOf((*MockCmdable)(nil).XClaimJustID), ctx, a)
}

// XDel mocks base method.
func (m *MockCmdable) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, stream}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XDel indicates an expected call of XDel.
func (mr *MockCmdableMockRecorder) XDel(ctx, stream interface{}, ids ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, stream}, ids...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XDel", reflect.TypeOf((*MockCmdable)(nil).XDel), varargs...)
}

// XGroupCreate mocks base method.
func (m *MockCmdable) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreate", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreate indicates an expected call of XGroupCreate.
func (mr *MockCmdableMockRecorder) XGroupCreate(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreate", reflect.TypeOf((*MockCmdable)(nil).XGroupCreate), ctx, stream, group, start)
}

// XGroupCreateConsumer mocks base method.
func (m *MockCmdable) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateConsumer", ctx, stream, group, consumer)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupCreateConsumer indicates an expected call of XGroupCreateConsumer.
func (mr *MockCmdableMockRecorder) XGroupCreateConsumer(ctx, stream, group, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateConsumer", reflect.TypeOf((*MockCmdable)(nil).XGroupCreateConsumer), ctx, stream, group, consumer)
}

// XGroupCreateMkStream mocks base method.
func (m *MockCmdable) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupCreateMkStream", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupCreateMkStream indicates an expected call of XGroupCreateMkStream.
func (mr *MockCmdableMockRecorder) XGroupCreateMkStream(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupCreateMkStream", reflect.TypeOf((*MockCmdable)(nil).XGroupCreateMkStream), ctx, stream, group, start)
}

// XGroupDelConsumer mocks base method.
func (m *MockCmdable) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDelConsumer", ctx, stream, group, consumer)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDelConsumer indicates an expected call of XGroupDelConsumer.
func (mr *MockCmdableMockRecorder) XGroupDelConsumer(ctx, stream, group, consumer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDelConsumer", reflect.TypeOf((*MockCmdable)(nil).XGroupDelConsumer), ctx, stream, group, consumer)
}

// XGroupDestroy mocks base method.
func (m *MockCmdable) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupDestroy", ctx, stream, group)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XGroupDestroy indicates an expected call of XGroupDestroy.
func (mr *MockCmdableMockRecorder) XGroupDestroy(ctx, stream, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupDestroy", reflect.TypeOf((*MockCmdable)(nil).XGroupDestroy), ctx, stream, group)
}

// XGroupSetID mocks base method.
func (m *MockCmdable) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XGroupSetID", ctx, stream, group, start)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// XGroupSetID indicates an expected call of XGroupSetID.
func (mr *MockCmdableMockRecorder) XGroupSetID(ctx, stream, group, start interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XGroupSetID", reflect.TypeOf((*MockCmdable)(nil).XGroupSetID), ctx, stream, group, start)
}

// XInfoConsumers mocks base method.
func (m *MockCmdable) XInfoConsumers(ctx context.Context, key, group string) *redis.XInfoConsumersCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoConsumers", ctx, key, group)
	ret0, _ := ret[0].(*redis.XInfoConsumersCmd)
	return ret0
}

// XInfoConsumers indicates an expected call of XInfoConsumers.
func (mr *MockCmdableMockRecorder) XInfoConsumers(ctx, key, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoConsumers", reflect.TypeOf((*MockCmdable)(nil).XInfoConsumers), ctx, key, group)
}

// XInfoGroups mocks base method.
func (m *MockCmdable) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoGroups", ctx, key)
	ret0, _ := ret[0].(*redis.XInfoGroupsCmd)
	return ret0
}

// XInfoGroups indicates an expected call of XInfoGroups.
func (mr *MockCmdableMockRecorder) XInfoGroups(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoGroups", reflect.TypeOf((*MockCmdable)(nil).XInfoGroups), ctx, key)
}

// XInfoStream mocks base method.
func (m *MockCmdable) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStream", ctx, key)
	ret0, _ := ret[0].(*redis.XInfoStreamCmd)
	return ret0
}

// XInfoStream indicates an expected call of XInfoStream.
func (mr *MockCmdableMockRecorder) XInfoStream(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStream", reflect.TypeOf((*MockCmdable)(nil).XInfoStream), ctx, key)
}

// XInfoStreamFull mocks base method.
func (m *MockCmdable) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XInfoStreamFull", ctx, key, count)
	ret0, _ := ret[0].(*redis.XInfoStreamFullCmd)
	return ret0
}

// XInfoStreamFull indicates an expected call of XInfoStreamFull.
func (mr *MockCmdableMockRecorder) XInfoStreamFull(ctx, key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XInfoStreamFull", reflect.TypeOf((*MockCmdable)(nil).XInfoStreamFull), ctx, key, count)
}

// XLen mocks base method.
func (m *MockCmdable) XLen(ctx context.Context, stream string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XLen", ctx, stream)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XLen indicates an expected call of XLen.
func (mr *MockCmdableMockRecorder) XLen(ctx, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XLen", reflect.TypeOf((*MockCmdable)(nil).XLen), ctx, stream)
}

// XPending mocks base method.
func (m *MockCmdable) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPending", ctx, stream, group)
	ret0, _ := ret[0].(*redis.XPendingCmd)
	return ret0
}

// XPending indicates an expected call of XPending.
func (mr *MockCmdableMockRecorder) XPending(ctx, stream, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPending", reflect.TypeOf((*MockCmdable)(nil).XPending), ctx, stream, group)
}

// XPendingExt mocks base method.
func (m *MockCmdable) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XPendingExt", ctx, a)
	ret0, _ := ret[0].(*redis.XPendingExtCmd)
	return ret0
}

// XPendingExt indicates an expected call of XPendingExt.
func (mr *MockCmdableMockRecorder) XPendingExt(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XPendingExt", reflect.TypeOf((*MockCmdable)(nil).XPendingExt), ctx, a)
}

// XRange mocks base method.
func (m *MockCmdable) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRange", ctx, stream, start, stop)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRange indicates an expected call of XRange.
func (mr *MockCmdableMockRecorder) XRange(ctx, stream, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRange", reflect.TypeOf((*MockCmdable)(nil).XRange), ctx, stream, start, stop)
}

// XRangeN mocks base method.
func (m *MockCmdable) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRangeN", ctx, stream, start, stop, count)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRangeN indicates an expected call of XRangeN.
func (mr *MockCmdableMockRecorder) XRangeN(ctx, stream, start, stop, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRangeN", reflect.TypeOf((*MockCmdable)(nil).XRangeN), ctx, stream, start, stop, count)
}

// XRead mocks base method.
func (m *MockCmdable) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRead", ctx, a)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XRead indicates an expected call of XRead.
func (mr *MockCmdableMockRecorder) XRead(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRead", reflect.TypeOf((*MockCmdable)(nil).XRead), ctx, a)
}

// XReadGroup mocks base method.
func (m *MockCmdable) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XReadGroup", ctx, a)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadGroup indicates an expected call of XReadGroup.
func (mr *MockCmdableMockRecorder) XReadGroup(ctx, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadGroup", reflect.TypeOf((*MockCmdable)(nil).XReadGroup), ctx, a)
}

// XReadStreams mocks base method.
func (m *MockCmdable) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range streams {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "XReadStreams", varargs...)
	ret0, _ := ret[0].(*redis.XStreamSliceCmd)
	return ret0
}

// XReadStreams indicates an expected call of XReadStreams.
func (mr *MockCmdableMockRecorder) XReadStreams(ctx interface{}, streams ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, streams...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XReadStreams", reflect.TypeOf((*MockCmdable)(nil).XReadStreams), varargs...)
}

// XRevRange mocks base method.
func (m *MockCmdable) XRevRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRange", ctx, stream, start, stop)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRange indicates an expected call of XRevRange.
func (mr *MockCmdableMockRecorder) XRevRange(ctx, stream, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRange", reflect.TypeOf((*MockCmdable)(nil).XRevRange), ctx, stream, start, stop)
}

// XRevRangeN mocks base method.
func (m *MockCmdable) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XRevRangeN", ctx, stream, start, stop, count)
	ret0, _ := ret[0].(*redis.XMessageSliceCmd)
	return ret0
}

// XRevRangeN indicates an expected call of XRevRangeN.
func (mr *MockCmdableMockRecorder) XRevRangeN(ctx, stream, start, stop, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XRevRangeN", reflect.TypeOf((*MockCmdable)(nil).XRevRangeN), ctx, stream, start, stop, count)
}

// XTrim mocks base method.
func (m *MockCmdable) XTrim(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrim", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrim indicates an expected call of XTrim.
func (mr *MockCmdableMockRecorder) XTrim(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrim", reflect.TypeOf((*MockCmdable)(nil).XTrim), ctx, key, maxLen)
}

// XTrimApprox mocks base method.
func (m *MockCmdable) XTrimApprox(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimApprox", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimApprox indicates an expected call of XTrimApprox.
func (mr *MockCmdableMockRecorder) XTrimApprox(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimApprox", reflect.TypeOf((*MockCmdable)(nil).XTrimApprox), ctx, key, maxLen)
}

// XTrimMaxLen mocks base method.
func (m *MockCmdable) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLen", ctx, key, maxLen)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLen indicates an expected call of XTrimMaxLen.
func (mr *MockCmdableMockRecorder) XTrimMaxLen(ctx, key, maxLen interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLen", reflect.TypeOf((*MockCmdable)(nil).XTrimMaxLen), ctx, key, maxLen)
}

// XTrimMaxLenApprox mocks base method.
func (m *MockCmdable) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMaxLenApprox", ctx, key, maxLen, limit)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMaxLenApprox indicates an expected call of XTrimMaxLenApprox.
func (mr *MockCmdableMockRecorder) XTrimMaxLenApprox(ctx, key, maxLen, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMaxLenApprox", reflect.TypeOf((*MockCmdable)(nil).XTrimMaxLenApprox), ctx, key, maxLen, limit)
}

// XTrimMinID mocks base method.
func (m *MockCmdable) XTrimMinID(ctx context.Context, key, minID string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinID", ctx, key, minID)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinID indicates an expected call of XTrimMinID.
func (mr *MockCmdableMockRecorder) XTrimMinID(ctx, key, minID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinID", reflect.TypeOf((*MockCmdable)(nil).XTrimMinID), ctx, key, minID)
}

// XTrimMinIDApprox mocks base method.
func (m *MockCmdable) XTrimMinIDApprox(ctx context.Context, key, minID string, limit int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XTrimMinIDApprox", ctx, key, minID, limit)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// XTrimMinIDApprox indicates an expected call of XTrimMinIDApprox.
func (mr *MockCmdableMockRecorder) XTrimMinIDApprox(ctx, key, minID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XTrimMinIDApprox", reflect.TypeOf((*MockCmdable)(nil).XTrimMinIDApprox), ctx, key, minID, limit)
}

// ZAdd mocks base method.
func (m *MockCmdable) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockCmdableMockRecorder) ZAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockCmdable)(nil).ZAdd), varargs...)
}

// ZAddArgs mocks base method.
func (m *MockCmdable) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddArgs", ctx, key, args)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddArgs indicates an expected call of ZAddArgs.
func (mr *MockCmdableMockRecorder) ZAddArgs(ctx, key, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddArgs", reflect.TypeOf((*MockCmdable)(nil).ZAddArgs), ctx, key, args)
}

// ZAddArgsIncr mocks base method.
func (m *MockCmdable) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZAddArgsIncr", ctx, key, args)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZAddArgsIncr indicates an expected call of ZAddArgsIncr.
func (mr *MockCmdableMockRecorder) ZAddArgsIncr(ctx, key, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddArgsIncr", reflect.TypeOf((*MockCmdable)(nil).ZAddArgsIncr), ctx, key, args)
}

// ZAddCh mocks base method.
func (m *MockCmdable) ZAddCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddCh indicates an expected call of ZAddCh.
func (mr *MockCmdableMockRecorder) ZAddCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddCh", reflect.TypeOf((*MockCmdable)(nil).ZAddCh), varargs...)
}

// ZAddNX mocks base method.
func (m *MockCmdable) ZAddNX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddNX indicates an expected call of ZAddNX.
func (mr *MockCmdableMockRecorder) ZAddNX(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNX", reflect.TypeOf((*MockCmdable)(nil).ZAddNX), varargs...)
}

// ZAddNXCh mocks base method.
func (m *MockCmdable) ZAddNXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddNXCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddNXCh indicates an expected call of ZAddNXCh.
func (mr *MockCmdableMockRecorder) ZAddNXCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddNXCh", reflect.TypeOf((*MockCmdable)(nil).ZAddNXCh), varargs...)
}

// ZAddXX mocks base method.
func (m *MockCmdable) ZAddXX(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddXX", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddXX indicates an expected call of ZAddXX.
func (mr *MockCmdableMockRecorder) ZAddXX(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddXX", reflect.TypeOf((*MockCmdable)(nil).ZAddXX), varargs...)
}

// ZAddXXCh mocks base method.
func (m *MockCmdable) ZAddXXCh(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAddXXCh", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAddXXCh indicates an expected call of ZAddXXCh.
func (mr *MockCmdableMockRecorder) ZAddXXCh(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAddXXCh", reflect.TypeOf((*MockCmdable)(nil).ZAddXXCh), varargs...)
}

// ZCard mocks base method.
func (m *MockCmdable) ZCard(ctx context.Context, key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", ctx, key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCard indicates an expected call of ZCard.
func (mr *MockCmdableMockRecorder) ZCard(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockCmdable)(nil).ZCard), ctx, key)
}

// ZCount mocks base method.
func (m *MockCmdable) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCount", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCount indicates an expected call of ZCount.
func (mr *MockCmdableMockRecorder) ZCount(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCount", reflect.TypeOf((*MockCmdable)(nil).ZCount), ctx, key, min, max)
}

// ZDiff mocks base method.
func (m *MockCmdable) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiff", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZDiff indicates an expected call of ZDiff.
func (mr *MockCmdableMockRecorder) ZDiff(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiff", reflect.TypeOf((*MockCmdable)(nil).ZDiff), varargs...)
}

// ZDiffStore mocks base method.
func (m *MockCmdable) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, destination}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffStore", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZDiffStore indicates an expected call of ZDiffStore.
func (mr *MockCmdableMockRecorder) ZDiffStore(ctx, destination interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, destination}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffStore", reflect.TypeOf((*MockCmdable)(nil).ZDiffStore), varargs...)
}

// ZDiffWithScores mocks base method.
func (m *MockCmdable) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZDiffWithScores", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZDiffWithScores indicates an expected call of ZDiffWithScores.
func (mr *MockCmdableMockRecorder) ZDiffWithScores(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZDiffWithScores", reflect.TypeOf((*MockCmdable)(nil).ZDiffWithScores), varargs...)
}

// ZIncr mocks base method.
func (m *MockCmdable) ZIncr(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncr", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncr indicates an expected call of ZIncr.
func (mr *MockCmdableMockRecorder) ZIncr(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncr", reflect.TypeOf((*MockCmdable)(nil).ZIncr), ctx, key, member)
}

// ZIncrBy mocks base method.
func (m *MockCmdable) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrBy", ctx, key, increment, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrBy indicates an expected call of ZIncrBy.
func (mr *MockCmdableMockRecorder) ZIncrBy(ctx, key, increment, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrBy", reflect.TypeOf((*MockCmdable)(nil).ZIncrBy), ctx, key, increment, member)
}

// ZIncrNX mocks base method.
func (m *MockCmdable) ZIncrNX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrNX", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrNX indicates an expected call of ZIncrNX.
func (mr *MockCmdableMockRecorder) ZIncrNX(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrNX", reflect.TypeOf((*MockCmdable)(nil).ZIncrNX), ctx, key, member)
}

// ZIncrXX mocks base method.
func (m *MockCmdable) ZIncrXX(ctx context.Context, key string, member *redis.Z) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZIncrXX", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZIncrXX indicates an expected call of ZIncrXX.
func (mr *MockCmdableMockRecorder) ZIncrXX(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZIncrXX", reflect.TypeOf((*MockCmdable)(nil).ZIncrXX), ctx, key, member)
}

// ZInter mocks base method.
func (m *MockCmdable) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInter", ctx, store)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZInter indicates an expected call of ZInter.
func (mr *MockCmdableMockRecorder) ZInter(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInter", reflect.TypeOf((*MockCmdable)(nil).ZInter), ctx, store)
}

// ZInterStore mocks base method.
func (m *MockCmdable) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterStore", ctx, destination, store)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZInterStore indicates an expected call of ZInterStore.
func (mr *MockCmdableMockRecorder) ZInterStore(ctx, destination, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterStore", reflect.TypeOf((*MockCmdable)(nil).ZInterStore), ctx, destination, store)
}

// ZInterWithScores mocks base method.
func (m *MockCmdable) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZInterWithScores", ctx, store)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZInterWithScores indicates an expected call of ZInterWithScores.
func (mr *MockCmdableMockRecorder) ZInterWithScores(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZInterWithScores", reflect.TypeOf((*MockCmdable)(nil).ZInterWithScores), ctx, store)
}

// ZLexCount mocks base method.
func (m *MockCmdable) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZLexCount", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZLexCount indicates an expected call of ZLexCount.
func (mr *MockCmdableMockRecorder) ZLexCount(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZLexCount", reflect.TypeOf((*MockCmdable)(nil).ZLexCount), ctx, key, min, max)
}

// ZMScore mocks base method.
func (m *MockCmdable) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZMScore", varargs...)
	ret0, _ := ret[0].(*redis.FloatSliceCmd)
	return ret0
}

// ZMScore indicates an expected call of ZMScore.
func (mr *MockCmdableMockRecorder) ZMScore(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZMScore", reflect.TypeOf((*MockCmdable)(nil).ZMScore), varargs...)
}

// ZPopMax mocks base method.
func (m *MockCmdable) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range count {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMax", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMax indicates an expected call of ZPopMax.
func (mr *MockCmdableMockRecorder) ZPopMax(ctx, key interface{}, count ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, count...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMax", reflect.TypeOf((*MockCmdable)(nil).ZPopMax), varargs...)
}

// ZPopMin mocks base method.
func (m *MockCmdable) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range count {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZPopMin", varargs...)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZPopMin indicates an expected call of ZPopMin.
func (mr *MockCmdableMockRecorder) ZPopMin(ctx, key interface{}, count ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, count...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZPopMin", reflect.TypeOf((*MockCmdable)(nil).ZPopMin), varargs...)
}

// ZRandMember mocks base method.
func (m *MockCmdable) ZRandMember(ctx context.Context, key string, count int, withScores bool) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRandMember", ctx, key, count, withScores)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRandMember indicates an expected call of ZRandMember.
func (mr *MockCmdableMockRecorder) ZRandMember(ctx, key, count, withScores interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRandMember", reflect.TypeOf((*MockCmdable)(nil).ZRandMember), ctx, key, count, withScores)
}

// ZRange mocks base method.
func (m *MockCmdable) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRange indicates an expected call of ZRange.
func (mr *MockCmdableMockRecorder) ZRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRange", reflect.TypeOf((*MockCmdable)(nil).ZRange), ctx, key, start, stop)
}

// ZRangeArgs mocks base method.
func (m *MockCmdable) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgs", ctx, z)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeArgs indicates an expected call of ZRangeArgs.
func (mr *MockCmdableMockRecorder) ZRangeArgs(ctx, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgs", reflect.TypeOf((*MockCmdable)(nil).ZRangeArgs), ctx, z)
}

// ZRangeArgsWithScores mocks base method.
func (m *MockCmdable) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeArgsWithScores", ctx, z)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeArgsWithScores indicates an expected call of ZRangeArgsWithScores.
func (mr *MockCmdableMockRecorder) ZRangeArgsWithScores(ctx, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeArgsWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeArgsWithScores), ctx, z)
}

// ZRangeByLex mocks base method.
func (m *MockCmdable) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByLex", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByLex indicates an expected call of ZRangeByLex.
func (mr *MockCmdableMockRecorder) ZRangeByLex(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRangeByLex), ctx, key, opt)
}

// ZRangeByScore mocks base method.
func (m *MockCmdable) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockCmdableMockRecorder) ZRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRangeByScore), ctx, key, opt)
}

// ZRangeByScoreWithScores mocks base method.
func (m *MockCmdable) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScoreWithScores", ctx, key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeByScoreWithScores indicates an expected call of ZRangeByScoreWithScores.
func (mr *MockCmdableMockRecorder) ZRangeByScoreWithScores(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScoreWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeByScoreWithScores), ctx, key, opt)
}

// ZRangeStore mocks base method.
func (m *MockCmdable) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeStore", ctx, dst, z)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRangeStore indicates an expected call of ZRangeStore.
func (mr *MockCmdableMockRecorder) ZRangeStore(ctx, dst, z interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeStore", reflect.TypeOf((*MockCmdable)(nil).ZRangeStore), ctx, dst, z)
}

// ZRangeWithScores mocks base method.
func (m *MockCmdable) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeWithScores", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeWithScores indicates an expected call of ZRangeWithScores.
func (mr *MockCmdableMockRecorder) ZRangeWithScores(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRangeWithScores), ctx, key, start, stop)
}

// ZRank mocks base method.
func (m *MockCmdable) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRank", ctx, key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRank indicates an expected call of ZRank.
func (mr *MockCmdableMockRecorder) ZRank(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRank", reflect.TypeOf((*MockCmdable)(nil).ZRank), ctx, key, member)
}

// ZRem mocks base method.
func (m *MockCmdable) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockCmdableMockRecorder) ZRem(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockCmdable)(nil).ZRem), varargs...)
}

// ZRemRangeByLex mocks base method.
func (m *MockCmdable) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByLex", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByLex indicates an expected call of ZRemRangeByLex.
func (mr *MockCmdableMockRecorder) ZRemRangeByLex(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByLex), ctx, key, min, max)
}

// ZRemRangeByRank mocks base method.
func (m *MockCmdable) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByRank", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByRank indicates an expected call of ZRemRangeByRank.
func (mr *MockCmdableMockRecorder) ZRemRangeByRank(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByRank", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByRank), ctx, key, start, stop)
}

// ZRemRangeByScore mocks base method.
func (m *MockCmdable) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRemRangeByScore", ctx, key, min, max)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRemRangeByScore indicates an expected call of ZRemRangeByScore.
func (mr *MockCmdableMockRecorder) ZRemRangeByScore(ctx, key, min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRemRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRemRangeByScore), ctx, key, min, max)
}

// ZRevRange mocks base method.
func (m *MockCmdable) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRange", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRange indicates an expected call of ZRevRange.
func (mr *MockCmdableMockRecorder) ZRevRange(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRange", reflect.TypeOf((*MockCmdable)(nil).ZRevRange), ctx, key, start, stop)
}

// ZRevRangeByLex mocks base method.
func (m *MockCmdable) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByLex", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByLex indicates an expected call of ZRevRangeByLex.
func (mr *MockCmdableMockRecorder) ZRevRangeByLex(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByLex", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByLex), ctx, key, opt)
}

// ZRevRangeByScore mocks base method.
func (m *MockCmdable) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScore", ctx, key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByScore indicates an expected call of ZRevRangeByScore.
func (mr *MockCmdableMockRecorder) ZRevRangeByScore(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScore", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByScore), ctx, key, opt)
}

// ZRevRangeByScoreWithScores mocks base method.
func (m *MockCmdable) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScoreWithScores", ctx, key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeByScoreWithScores indicates an expected call of ZRevRangeByScoreWithScores.
func (mr *MockCmdableMockRecorder) ZRevRangeByScoreWithScores(ctx, key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScoreWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeByScoreWithScores), ctx, key, opt)
}

// ZRevRangeWithScores mocks base method.
func (m *MockCmdable) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeWithScores", ctx, key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeWithScores indicates an expected call of ZRevRangeWithScores.
func (mr *MockCmdableMockRecorder) ZRevRangeWithScores(ctx, key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeWithScores", reflect.TypeOf((*MockCmdable)(nil).ZRevRangeWithScores), ctx, key, start, stop)
}

// ZRevRank mocks base method.
func (m *MockCmdable) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRank", ctx, key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRevRank indicates an expected call of ZRevRank.
func (mr *MockCmdableMockRecorder) ZRevRank(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRank", reflect.TypeOf((*MockCmdable)(nil).ZRevRank), ctx, key, member)
}

// ZScan mocks base method.
func (m *MockCmdable) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScan", ctx, key, cursor, match, count)
	ret0, _ := ret[0].(*redis.ScanCmd)
	return ret0
}

// ZScan indicates an expected call of ZScan.
func (mr *MockCmdableMockRecorder) ZScan(ctx, key, cursor, match, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScan", reflect.TypeOf((*MockCmdable)(nil).ZScan), ctx, key, cursor, match, count)
}

// ZScore mocks base method.
func (m *MockCmdable) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScore", ctx, key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZScore indicates an expected call of ZScore.
func (mr *MockCmdableMockRecorder) ZScore(ctx, key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScore", reflect.TypeOf((*MockCmdable)(nil).ZScore), ctx, key, member)
}

// ZUnion mocks base method.
func (m *MockCmdable) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnion", ctx, store)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZUnion indicates an expected call of ZUnion.
func (mr *MockCmdableMockRecorder) ZUnion(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnion", reflect.TypeOf((*MockCmdable)(nil).ZUnion), ctx, store)
}

// ZUnionStore mocks base method.
func (m *MockCmdable) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionStore", ctx, dest, store)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZUnionStore indicates an expected call of ZUnionStore.
func (mr *MockCmdableMockRecorder) ZUnionStore(ctx, dest, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionStore", reflect.TypeOf((*MockCmdable)(nil).ZUnionStore), ctx, dest, store)
}

// ZUnionWithScores mocks base method.
func (m *MockCmdable) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZUnionWithScores", ctx, store)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZUnionWithScores indicates an expected call of ZUnionWithScores.
func (mr *MockCmdableMockRecorder) ZUnionWithScores(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZUnionWithScores", reflect.TypeOf((*MockCmdable)(nil).ZUnionWithScores), ctx, store)
}
