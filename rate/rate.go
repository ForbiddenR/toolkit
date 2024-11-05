package rate

import (
	"sync"
	"time"
)

type Limiter struct {
	duration time.Duration
	actTime  time.Time
	lock     sync.Locker
}

func NewLimiter(duration time.Duration) *Limiter {
	limiter := &Limiter{
		duration: duration,
		lock:     &sync.Mutex{},
	}
	limiter.actTime = time.Now()
	return limiter
}

func (l *Limiter) Allow() (time.Duration, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now()
	if diff := l.actTime.Sub(now); diff > 0 {
		return diff, false
	}
	l.actTime = now.Add(l.duration)
	return time.Duration(0), true
}

func (l *Limiter) Reach() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now()
	if diff := l.actTime.Sub(now); diff > 0 {
		return false
	}
	l.actTime = now.Add(l.duration)
	return true
}
