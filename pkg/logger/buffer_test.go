package logger

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuffer_Write(t *testing.T) {
	buf := newBuffer()
	size, err := buf.Write([]byte("hello world"))
	assert.Equal(t, 11, size)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())
}

func TestBuffer_WriteString(t *testing.T) {
	buf := newBuffer()
	size, err := buf.WriteString("hello world")
	assert.Equal(t, 11, size)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())
}

func TestBuffer_WriteTo(t *testing.T) {
	buf := newBuffer()
	_, _ = buf.WriteString("hello world")
	buf2 := bytes.NewBuffer(nil)
	size, err := buf.WriteTo(buf2)
	assert.Equal(t, int64(11), size)
	assert.NoError(t, err)
	assert.Equal(t, "hello world", buf.String())
	assert.Equal(t, "hello world", buf2.String())
}

func TestBuffer_AppendByte(t *testing.T) {
	var p byte = 65
	buf := newBuffer()
	buf.AppendByte(p)
	assert.Equal(t, string(p), buf.String())
}

func TestBuffer_AppendBytes(t *testing.T) {
	buf := newBuffer()
	buf.AppendBytes([]byte("hello world"))
	assert.Equal(t, "hello world", buf.String())
}

func TestBuffer_AppendString(t *testing.T) {
	buf := newBuffer()
	buf.AppendString("hello world")
	assert.Equal(t, "hello world", buf.String())
}

func TestBuffer_AppendQuote(t *testing.T) {
	buf := newBuffer()
	buf.AppendQuote("hello world")
	assert.Equal(t, `"hello world"`, buf.String())
}

func TestBuffer_AppendInt(t *testing.T) {
	buf := newBuffer()
	buf.AppendInt(123)
	assert.Equal(t, "123", buf.String())
}

func TestBuffer_AppendUint(t *testing.T) {
	buf := newBuffer()
	buf.AppendUint(123)
	assert.Equal(t, "123", buf.String())
}

func TestBuffer_AppendFloat32(t *testing.T) {
	buf := newBuffer()
	buf.AppendFloat32(1.1)
	assert.Equal(t, "1.1", buf.String())
}

func TestBuffer_AppendFloat64(t *testing.T) {
	buf := newBuffer()
	buf.AppendFloat64(1.1)
	assert.Equal(t, "1.1", buf.String())
}

func TestBuffer_AppendBool(t *testing.T) {
	buf := newBuffer()
	buf.AppendBool(true)
	buf.AppendBool(false)
	assert.Equal(t, "truefalse", buf.String())
}

func TestBuffer_AppendTimeFormat(t *testing.T) {
	now := time.Now()
	buf := newBuffer()
	buf.AppendTimeFormat(now, time.RFC3339)
	assert.Equal(t, now.Format(time.RFC3339), buf.String())
}

func TestBuffer_Replace(t *testing.T) {
	var p = []byte("!")
	buf := newBuffer()
	buf.AppendBytes([]byte("hello world?"))
	buf.Replace(11, p[0])
	assert.Equal(t, "hello world!", buf.String())
}

func TestBuffer_ReplaceOut(t *testing.T) {
	var p = []byte("!")
	buf := newBuffer()
	buf.AppendBytes([]byte("hello world?"))
	buf.Replace(21, p[0])
	assert.Equal(t, "hello world?", buf.String())
}

func TestBuffer_Len(t *testing.T) {
	buf := newBuffer()
	buf.AppendString("hello world")
	assert.Equal(t, 11, buf.Len())
}

func TestBuffer_Cap(t *testing.T) {
	buf := newBuffer()
	buf.AppendString("hello world")
	assert.Equal(t, bufferSize, buf.Cap())
}

func TestBuffer_String(t *testing.T) {
	buf := newBuffer()
	buf.AppendString("hello world")
	assert.Equal(t, "hello world", buf.String())
}

func TestBuffer_Reset(t *testing.T) {
	buf := newBuffer()
	buf.AppendString("hello world")
	buf.Reset()
	buf.AppendString("hello world 2")
	assert.Equal(t, "hello world 2", buf.String())
}
