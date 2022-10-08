package oss_test

import (
	"runtime/debug"
	"testing"

	"github.com/xuender/oils/oss"
)

func TestGetMod(t *testing.T) {
	t.Parallel()

	if mod := oss.GetMod("test"); mod == nil {
		t.Errorf("GetMod() return= %v, wantErr %v", mod, nil)
	}
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
