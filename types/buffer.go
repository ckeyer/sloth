package types

import (
	"bytes"
	"fmt"
	"sync"
)

// WriteHandler:
type WriteHandler interface{}

// Buffer:
type Buffer struct {
	sync.Mutex
	buf      *bytes.Buffer
	times    int
	SyncRead bool
}

// NewBuf
func NewBuf() (buf *Buffer) {
	buf = &Buffer{
		buf: &bytes.Buffer{},
	}
	return
}

// Write:
func (b *Buffer) Write(data []byte) (int, error) {
	b.Lock()
	defer b.Unlock()
	fmt.Print("ing...", string(data))
	b.times++
	return b.buf.Write(data)
}

// Clear:
func (b *Buffer) Clear() {
	b.Lock()
	defer b.Unlock()
	b.buf.Reset()
	b.times = 0
}

// Over:
func (b *Buffer) Over() {
	fmt.Println(b.String())
	fmt.Println("length: ", b.Len())
}

// GetInputTimes:
func (b *Buffer) GetInputTimes() int {
	return b.times
}

// Len:
func (b *Buffer) Len() int {
	return b.buf.Len()
}

// String:
func (b *Buffer) String() string {
	return b.buf.String()
}
