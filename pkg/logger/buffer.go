package logger

import (
	"io"
	"strconv"
	"sync"
	"time"
)

const (
	bufferSize        = 1024
	poolMaxBufferSize = 16 << 10
)

type bufferPool struct {
	pool sync.Pool
}

func newBufferPool() *bufferPool {
	return &bufferPool{
		pool: sync.Pool{
			New: func() any {
				return newBuffer()
			},
		},
	}
}

func (p *bufferPool) Acquire() *Buffer {
	return p.pool.Get().(*Buffer)
}

func (p *bufferPool) Free(b *Buffer) {
	if cap(b.buf) <= poolMaxBufferSize {
		b.Reset()
		p.pool.Put(b)
	}
}

type Buffer struct {
	buf []byte
}

func newBuffer() *Buffer {
	return &Buffer{buf: make([]byte, 0, bufferSize)}
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.buf = append(b.buf, p...)
	return len(p), nil
}

func (b *Buffer) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}

func (b *Buffer) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(b.buf)
	return int64(n), err
}

func (b *Buffer) AppendByte(p byte) {
	b.buf = append(b.buf, p)
}

func (b *Buffer) AppendBytes(p []byte) {
	b.buf = append(b.buf, p...)
}

func (b *Buffer) AppendString(s string) {
	b.buf = append(b.buf, s...)
}

func (b *Buffer) AppendQuote(s string) {
	b.buf = strconv.AppendQuote(b.buf, s)
}

func (b *Buffer) AppendInt(i int64) {
	b.buf = strconv.AppendInt(b.buf, i, 10) //nolint:mnd
}

func (b *Buffer) AppendUint(i uint64) {
	b.buf = strconv.AppendUint(b.buf, i, 10) //nolint:mnd
}

func (b *Buffer) AppendFloat32(f float32) {
	b.AppendFloat(float64(f), 32) //nolint:mnd
}

func (b *Buffer) AppendFloat64(f float64) {
	b.AppendFloat(f, 64) //nolint:mnd
}

func (b *Buffer) AppendFloat(f float64, bitSize int) {
	b.buf = strconv.AppendFloat(b.buf, f, 'f', -1, bitSize)
}

func (b *Buffer) AppendBool(v bool) {
	b.buf = strconv.AppendBool(b.buf, v)
}

func (b *Buffer) AppendTimeFormat(t time.Time, layout string) {
	b.buf = t.AppendFormat(b.buf, layout)
}

func (b *Buffer) Replace(i int, p byte) {
	if i < 0 || i >= b.Len() {
		return
	}
	b.buf[i] = p
}

func (b *Buffer) Len() int {
	return len(b.buf)
}

func (b *Buffer) Cap() int {
	return cap(b.buf)
}

func (b *Buffer) String() string {
	return string(b.buf)
}

func (b *Buffer) Reset() {
	b.buf = b.buf[:0]
}
