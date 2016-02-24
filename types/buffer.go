package types

import (
	"bytes"
	"fmt"
	"sync"
)

type WriteHandler interface{}

type Buffer struct {
	sync.Mutex
	buf      *bytes.Buffer
	times    int
	SyncRead bool
}

func NewBuf() (buf *Buffer) {
	buf = &Buffer{
		buf: &bytes.Buffer{},
	}
	return
}

func (b *Buffer) Write(data []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	fmt.Print("ing...", string(data))
	b.times++
	return b.buf.Write(data)
}

func (b *Buffer) Clear() {
	b.Lock()
	defer b.Unlock()
	b.buf.Reset()
	b.times = 0
}

func (b *Buffer) Over() {
	fmt.Println(b.String())
	fmt.Println("length: ", b.Len())
}

func (b *Buffer) GetInputTimes() int {
	return b.times
}

func (b *Buffer) Len() int {
	return b.buf.Len()
}

func (b *Buffer) String() string {
	return b.buf.String()
}
