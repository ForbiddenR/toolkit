package buffer

import (
	"sync"

	"github.com/ForbiddenR/toolkit/pool"
)

type Pool struct {
	// p *sync.Pool
	p *pool.Pool[*Buffer]
}

func NewPool() Pool {
	return Pool{
		p: pool.New()
	}
}

func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.Reset()
	buf.pool = p
	return buf
}

func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}
