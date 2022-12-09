package oss_test

import (
	"runtime/debug"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/oss"
)

// nolint: paralleltest
func TestGetMod(t *testing.T) {
	assert.NotNil(t, oss.GetMod("test"))

	patches := gomonkey.ApplyFuncReturn(debug.ReadBuildInfo, nil, false)
	defer patches.Reset()

	assert.Nil(t, oss.GetMod("test"))
}

// nolint: exhaustruct
func TestGetModByInfo(t *testing.T) {
	t.Parallel()

	module := debug.Module{Replace: &debug.Module{}}
	info := &debug.BuildInfo{Main: module, Deps: []*debug.Module{{
		Path:    "xuender/gosign",
		Replace: &debug.Module{},
	}}}

	if mod := oss.GetModByInfo(info, "gosign"); mod == nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, module)
	}
}
