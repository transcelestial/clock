package clocktest

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

func TestClockWithTicker(t *testing.T) {
	ticker := NewTicker()
	c := New(WithTicker(ticker))
	ct := c.NewTicker(time.Second)
	assert.Equal(t, ticker, ct)
	// no need to test the ticker here
}

func TestClockWithoutTicker(t *testing.T) {
	c := New()
	assert.Nil(t, c.NewTicker(time.Second))
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

func TestClockWithSleepr(t *testing.T) {
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
