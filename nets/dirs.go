package nets

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xuender/oils/base"
)

type Dirs []string

func (p Dirs) Open(name string) (http.File, error) {
	if len(p) == 0 {
		defaultDir := http.Dir(".")

		return defaultDir.Open(name)
	}

	for _, dir := range p {
		base := "/" + filepath.Base(dir)
		if strings.HasPrefix(name, base) {
			return http.Dir(dir).Open(name[len(base):])
		}
	}

	if name == "" || name == "." || name == "/" {
		return p, nil
	}

	return nil, fs.ErrNotExist
}

func (p Dirs) Readdir(count int) ([]fs.FileInfo, error) {
	fileInfos := make([]fs.FileInfo, len(p))

	for index, dir := range p {
		fileInfos[index] = base.Must1(os.Stat(dir))
	}

	return fileInfos, nil
}

func (p Dirs) Stat() (fs.FileInfo, error)                   { return p, nil }
func (p Dirs) Name() string                                 { return strings.Join(p, ",") }
func (p Dirs) Size() int64                                  { return 0 }
func (p Dirs) Mode() fs.FileMode                            { return fs.ModeDir }
func (p Dirs) ModTime() time.Time                           { return time.Now() }
func (p Dirs) IsDir() bool                                  { return true }
func (p Dirs) Sys() any                                     { return p }
func (p Dirs) Close() error                                 { return nil }
func (p Dirs) Read(data []byte) (n int, err error)          { return 0, nil }
func (p Dirs) Seek(offset int64, whence int) (int64, error) { return offset, nil }
