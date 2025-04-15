package cache

import "sync"

type Cache[K comparable, V any] interface {
	Set(K, V)
	Get(K) (V, bool)
}

type defaultCache[K comparable, V any] struct {
	cache sync.Map
}

func NewDefaultCache[K comparable, V any]() Cache[K, V] {
	return &defaultCache[K, V]{}
}

func (c *defaultCache[K, V]) Set(k K, v V) {
	c.cache.Store(k, v)
}

func (c *defaultCache[K, V]) Get(k K) (V, bool) {
	v, ok := c.cache.Load(k)
	if !ok {
		var zero V
		return zero, false
	}
	return v.(V), true
}
