package rate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	limiter := NewLimiter(time.Second)
	duration, ok := limiter.Allow()
	assert.Equal(t, time.Duration(0), duration)
	assert.Equal(t, true, ok)

	duration, ok = limiter.Allow()
	assert.Equal(t, false, ok)
	assert.LessOrEqual(t, duration, 1*time.Second)
	t.Log(duration.String())
	time.Sleep(duration)
	duration, ok = limiter.Allow()
	assert.Equal(t, time.Duration(0), duration)
	assert.Equal(t, true, ok)
}
