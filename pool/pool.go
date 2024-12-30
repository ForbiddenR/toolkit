package pool

import "sync"

type Pool[T any] struct {
	p *sync.Pool
}

func New[T any](new func() T) *Pool[T] {
	return &Pool[T]{
		p: &sync.Pool{
			New: func() interface{} {
				return new()
			},
		},
	}
}

func (p Pool[t]) Get() t {
	return p.p.Get().(t)
}

func (p Pool[t]) Put(r t) {
	p.p.Put(r)
}
