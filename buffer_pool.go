package gintool

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

func AcquireBuffer() *bytes.Buffer {
	v := bufferPool.Get()
	if v == nil {
		return &bytes.Buffer{}
	}
	return v.(*bytes.Buffer)
}

func ReleaseBuffer(buffer *bytes.Buffer) {
	if buffer != nil {
		buffer.Reset()
		bufferPool.Put(buffer)
	}
}
