package clocktest

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/transcelestial/clock"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

func TestClockWithTicker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	t1id := 0
	t2id := 1
	tt1 := NewTicker()
	tt2 := NewTicker()
	c := New(WithTicker(t1id, tt1), WithTicker(t2id, tt2))
	tk1 := c.NewTicker(time.Second, clock.TickerWithID(t1id))
	tk2 := c.NewTicker(time.Second, clock.TickerWithID(t2id))

	var wg sync.WaitGroup
	var t1Ticks []time.Time
	var t2Ticks []time.Time

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case tick := <-tk1.C():
				t1Ticks = append(t1Ticks, tick)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case tick := <-tk2.C():
				t2Ticks = append(t2Ticks, tick)
			}
		}
	}()

	t0 := time.Now()
	t1 := t0.Add(time.Second)
	t2 := t1.Add(time.Second)
	t3 := t2.Add(time.Second)
	t4 := t3.Add(time.Second)

	tt1.Next(ctx, t1, t2)
	tk1.Stop()
	tt2.Next(ctx, t3, t4)

	cancel()
	wg.Wait()

	if assert.Len(t, t1Ticks, 2) {
		assert.Equal(t, []time.Time{t1, t2}, t1Ticks)
	}

	if assert.Len(t, t2Ticks, 2) {
		assert.Equal(t, []time.Time{t3, t4}, t2Ticks)
	}
}

func TestClockWithoutTicker(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("expected a panic, but got nothing")
		}
	}()
	c := New()
	assert.Nil(t, c.NewTicker(time.Second))
}

func TestClockWithDupTickers(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("expected a panic, but got nothing")
		}
	}()
	c := New()
	opt := clock.TickerWithID(0)
	c.NewTicker(time.Second, opt)
	c.NewTicker(time.Second, opt)
}

func TestClockLoadBadTicker(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("expected a panic, but got nothing")
		}
	}()
	t1 := NewTicker()
	t2 := NewTicker()
	New(WithTicker(0, t1), WithTicker(0, t2))
}

func TestClockWithTimeQueue(t *testing.T) {
	defer leaktest.Check(t)()

	q := NewTimeQueue(2)
	c := New(WithTimeQueue(q))

	t0 := time.Now()
	t1 := time.Now()

	go func() {
		q.Push(t0, t1)
	}()

	t2 := c.Now()
	t3 := c.Now()

	assert.Equal(t, t0, t2)
	assert.Equal(t, t3, t1)
}

func TestClockWithEmptyTimeQueue(t *testing.T) {
	defer leaktest.Check(t)()

	q := NewTimeQueue(1)
	c := New(WithTimeQueue(q))

	ch := make(chan time.Time, 1)
	go func() {
		v := c.Now()
		ch <- v
	}()

	select {
	case <-ch:
		t.Errorf("unexpected time in queue")
		return
	case <-time.After(100 * time.Millisecond):
		// ok
	}

	t0 := time.Now()
	q.Push(t0)
	t1 := <-ch
	assert.Equal(t, t0, t1)
}

func TestClockWithSleeper(t *testing.T) {
	ctx := context.Background()

	s := NewSleeper()
	c := New(WithSleeper(s))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c.Sleep(time.Second)
		wg.Done()
	}()

	s.WakeUp(ctx)
	wg.Wait()
}
