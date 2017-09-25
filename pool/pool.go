package pool

import (
	"bytes"
	"sync"
)

// pattern of simple
func pool1(in string) (out []byte) {
	buf := &bytes.Buffer{}
	for i := 0; i < 100; i++ {
		buf.WriteString(in)
	}
	out = buf.Bytes()
	return
}

var globalBuf = &bytes.Buffer{}
var globalMutex = &sync.Mutex{}

// pattern of sharing buffer and using lock
func pool2(in string) (out []byte) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalBuf.Reset()
	for i := 0; i < 100; i++ {
		globalBuf.WriteString(in)
	}
	out = globalBuf.Bytes()
	return
}

var globalPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// pattern of using sync.Pool
func pool3(in string) (out []byte) {
	buf := globalPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		globalPool.Put(buf)
	}()
	for i := 0; i < 100; i++ {
		buf.WriteString(in)
	}

	out = buf.Bytes()
	return
}
