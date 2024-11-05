package id

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdGenerator(t *testing.T) {
	tig := NewTransactionIdGenerator()
	group := &sync.WaitGroup{}
	group.Add(100)
	for i := 0; i < 100; i++ {
		go func(group *sync.WaitGroup) {
			for i := 0; i < 100; i++ {
				id := tig.Get("32010600019236", "1")
				t.Log(id)
				assert.Equal(t, 32, len(id))
			}
			group.Done()
		}(group)
	}
	group.Wait()
}
