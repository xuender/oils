package oss

import (
	"runtime/debug"
	"strings"
)

func GetMod(name string) *debug.Module {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}

	return GetModByInfo(info, name)
}

func GetModByInfo(info *debug.BuildInfo, name string) *debug.Module {
	ret := &info.Main

	if ret.Replace != nil {
		ret = ret.Replace
	}

	if ret.Path == "" {
		for _, m := range info.Deps {
			if strings.HasSuffix(m.Path, name) {
				ret = m

				break
			}
		}
	}

	if ret.Replace != nil {
		return ret.Replace
	}

	return ret
}
