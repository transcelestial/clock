package clocktest

import (
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/stretchr/testify/assert"
)

func TestTimeQueue(t *testing.T) {
	defer leaktest.Check(t)()

	q := NewTimeQueue(2)
	t0 := time.Now()
	t1 := t0.Add(time.Second)
	q.Push(t0, t1)

	t2 := q.Drain()
	t3 := q.Drain()

	assert.Equal(t, t0, t2)
	assert.Equal(t, t1, t3)
}

func TestTimeQueueEmpty(t *testing.T) {
	defer leaktest.Check(t)()

	q := NewTimeQueue(1)

	c := make(chan time.Time, 1)
	go func() {
		v := q.Drain()
		c <- v
	}()

	select {
	case <-c:
		t.Errorf("unexpected time in queue")
		return
	case <-time.After(100 * time.Millisecond):
		// ok
	}

	t0 := time.Now()
	q.Push(t0)
	t1 := <-c
	assert.Equal(t, t0, t1)
}
