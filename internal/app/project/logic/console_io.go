package logic

import (
	"bytes"
	"fmt"
)

// 实现 Writer 接口
type Writer struct {
	buffer bytes.Buffer
}

func (w *Writer) Write(p []byte) (n int, err error) {
	fmt.Println("写入内容: ", string(p))
	return w.buffer.Write(p)
}

func (w *Writer) String() string {
	return w.buffer.String()
}

func NewWriter() *Writer {
	return &Writer{}
}
