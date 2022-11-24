package nets_test

import (
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/base"
	"github.com/xuender/oils/nets"
)

func TestDirs_Open(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	file, err := dirs.Open("/nets")

	assert.Nil(t, err)
	assert.NotNil(t, file)

	file, err = dirs.Open("")

	assert.Nil(t, err)
	assert.NotNil(t, file)

	_, err = dirs.Open("/sss")
	assert.NotNil(t, err)

	dirs = nets.Dirs([]string{})
	file, err = dirs.Open("")
	assert.Nil(t, err)
	assert.NotNil(t, file)
}

func TestDirs_Readdir(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.Equal(t, 2, len(base.Must1(dirs.Readdir(0))))
}

func TestDirs_Stat(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.True(t, base.Must1(dirs.Stat()).IsDir())
}

func TestDirs_Name(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.Equal(t, "../nets,../oss", dirs.Name())
}

func TestDirs_Size(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.Equal(t, int64(0), dirs.Size())
}

func TestDirs_Mode(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.Equal(t, fs.ModeDir, dirs.Mode())
}

func TestDirs_ModTime(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.LessOrEqual(t, dirs.ModTime().Unix(), time.Now().Unix())
}

func TestDirs_IsDir(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.True(t, dirs.IsDir())
}

func TestDirs_Sys(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.NotNil(t, dirs.Sys())
}

func TestDirs_Close(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	assert.Nil(t, dirs.Close())
}

func TestDirs_Read(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	_, err := dirs.Read([]byte(""))
	assert.Nil(t, err)
}

func TestDirs_Seek(t *testing.T) {
	t.Parallel()

	dirs := nets.Dirs([]string{"../nets", "../oss"})
	_, err := dirs.Seek(0, 0)
	assert.Nil(t, err)
}
