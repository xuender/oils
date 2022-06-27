package ios

import (
	"io"
	"os"

	"github.com/xuender/oils/oss"
)

// NewWriter 新建Writer，file存在Seek末尾,不存在创建.
func NewWriter(file string) (io.WriteCloser, error) {
	if oss.Exist(file) {
		writer, err := os.OpenFile(file, os.O_WRONLY, oss.DefaultFileMode)
		if err != nil {
			return writer, err
		}

		_, err = writer.Seek(0, os.SEEK_END)

		return writer, err
	}

	oss.MkdirParent(file)

	return os.Create(file)
}
