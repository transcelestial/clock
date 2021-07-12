package clocktest

import (
	"context"
	"sync"
	"time"

	"github.com/transcelestial/clock"
)

// NewTicker creates a new test ticker.
func NewTicker() *Ticker {
	t := &Ticker{
		c:       make(chan time.Time),
		running: true,
	}
	t.wg.Add(1)
	return t
}

type Ticker struct {
	c chan time.Time

	wg      sync.WaitGroup
	mux     sync.RWMutex
	running bool
}

// C returns the ticker channel
func (ticker *Ticker) C() <-chan time.Time {
	return ticker.c
}

// Reset sets the ticker as running
func (ticker *Ticker) Reset(d time.Duration) {
	ticker.mux.Lock()
	defer ticker.mux.Unlock()
	if !ticker.running {
		ticker.wg.Add(1)
	}
	ticker.running = true
}

// Stop sets the ticker as stopped
func (ticker *Ticker) Stop() {
	ticker.mux.Lock()
	defer ticker.mux.Unlock()
	if ticker.running {
		defer ticker.wg.Done()
	}
	ticker.running = false
}

// Next sends ticks on the C channel.
// It blocks until C receives the ticks. Use the ctx to cancel the tick.
// If Stop was called, the tick is dropped.
func (ticker *Ticker) Next(ctx context.Context, times ...time.Time) {
	ticker.mux.RLock()
	defer ticker.mux.RUnlock()
	if !ticker.running {
		return
	}
	for _, t := range times {
		select {
		case ticker.c <- t:
		case <-ctx.Done():
			return
		}
	}
}

// Done blocks the thread until the ticker is stopped.
// Subsequent calls to Done() after stop will not block.
func (ticker *Ticker) Done() {
	ticker.wg.Wait()
}

var _ clock.Ticker = &Ticker{}
