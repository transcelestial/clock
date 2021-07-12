package clocktest

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

func TestTicker(t *testing.T) {
	defer leaktest.Check(t)()

	ticker := NewTicker()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var ticks []time.Time
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case tick := <-ticker.C():
				ticks = append(ticks, tick)
			}
		}
	}()

	t0 := time.Now()
	t1 := t0.Add(time.Second)
	t2 := t1.Add(time.Second)
	t3 := t2.Add(time.Second)
	t4 := t3.Add(time.Second)

	ticker.Next(ctx, t0, t1, t2)
	ticker.Stop()
	ticker.Next(ctx, t3)
	ticker.Reset(time.Second)
	ticker.Next(ctx, t4)

	cancel()
	wg.Wait()

	if assert.Len(t, ticks, 4) {
		assert.Equal(t, []time.Time{t0, t1, t2, t4}, ticks)
	}
}

func TestTickerDone(t *testing.T) {
	defer leaktest.Check(t)()

	ticker := NewTicker()

	go func() {
		time.Sleep(time.Second)
		ticker.Stop()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ticker.Done()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		ticker.Done()
		wg.Done()
	}()

	wg.Wait()

	ticker.Done()
}

func TestTickerManyStop(t *testing.T) {
	defer leaktest.Check(t)()

	ticker := NewTicker()

	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()

	ticker.Stop()
	ticker.Stop()
	ticker.Done()
}

func TestTickerManyReset(t *testing.T) {
	defer leaktest.Check(t)()

	ticker := NewTicker()

	defer func() {
		err := recover()
		if err != nil {
			t.Fatal(err)
		}
	}()

	ticker.Reset(time.Second)
	ticker.Reset(time.Second)
	ticker.Stop()
	ticker.Done()
}
