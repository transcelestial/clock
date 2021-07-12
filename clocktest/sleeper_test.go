package clocktest

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSleeper(t *testing.T) {
	ctx := context.Background()
	s := NewSleeper()

	var wg sync.WaitGroup
	var mux sync.Mutex
	var diffs []time.Duration

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			now := time.Now()
			s.Sleep()
			mux.Lock()
			defer mux.Unlock()
			diffs = append(diffs, time.Since(now))
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		s.WakeUp(ctx)
	}

	wg.Wait()

	if assert.Len(t, diffs, 10) {
		for _, d := range diffs {
			assert.True(t, d >= 100*time.Millisecond)
		}
	}
}
func TestSleeperCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := NewSleeper()
	time.AfterFunc(100*time.Millisecond, cancel)
	s.WakeUp(ctx)
}
