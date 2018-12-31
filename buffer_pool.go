package gintool

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

func AcquireBuffer() *bytes.Buffer {
	if v := bufferPool.Get(); v != nil {
		return v.(*bytes.Buffer)
	}
	return &bytes.Buffer{}
}

func ReleaseBuffer(buffer *bytes.Buffer) {
	if buffer != nil {
		buffer.Reset()
		bufferPool.Put(buffer)
	}
}
