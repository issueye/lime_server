package logic

import (
	"bytes"
	"fmt"
	"lime/internal/app/project/model"
)

// 实现 Writer 接口
type Writer struct {
	buffer  bytes.Buffer
	Version model.VersionInfo
}

func (w *Writer) Write(p []byte) (n int, err error) {
	fmt.Println("写入内容: ", string(p))
	if w.Version.ID != 0 {
		// SendMessage(w.Version, string(p))
	}

	return w.buffer.Write(p)
}

func (w *Writer) String() string {
	return w.buffer.String()
}

func NewWriter(version model.VersionInfo) *Writer {
	return &Writer{
		Version: version,
	}
}
