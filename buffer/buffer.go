package buffer

import (
	"strconv"
	"unicode/utf8"
)

const _size = 1024

type Buffer struct {
	bs   []byte
	pool Pool
}

func (b *Buffer) Bytes() []byte {
	return b.bs
}

func (b *Buffer) String() string {
	return string(b.bs)
}

func (b *Buffer) Write(bs []byte) (int, error) {
	b.bs = append(b.bs, bs...)
	return len(bs), nil
}

func (b *Buffer) WriteByte(v byte) error {
	b.bs = append(b.bs, v)
	return nil
}

func (b *Buffer) WriteRune(r rune) (int, error) {
	n := len(b.bs)
	b.bs = utf8.AppendRune(b.bs, r)
	return len(b.bs) - n, nil
}

func (b *Buffer) WriteString(s string) (int, error) {
	b.bs = append(b.bs, s...)
	return len(s), nil
}

func (b *Buffer) WriteInt(i int64) error {
	b.bs = strconv.AppendInt(b.bs, i, 10)
	return nil
}

func (b *Buffer) Reset() {
	b.bs = b.bs[:0]
}

func (b *Buffer) Free() {
	b.pool.put(b)
}
