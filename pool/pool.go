package pool

import "sync"

type Resource interface {
	Reset()
}

type Pool[t Resource] struct {
	p *sync.Pool
}

func NewPool[t Resource](new func() t) *Pool[t] {
	return &Pool[t]{
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
	r.Reset()
	p.p.Put(r)
}
