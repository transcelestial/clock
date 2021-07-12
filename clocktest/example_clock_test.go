package clocktest_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/transcelestial/clock"
	"github.com/transcelestial/clock/clocktest"
)

func ExampleClock_Now() {
	// create a time queue w/ size of 2
	q := clocktest.NewTimeQueue(2)
	// and a test clock that uses the time queue
	c := clocktest.New(clocktest.WithTimeQueue(q))

	// inject the Clock interface as dependency into some func
	// you want to test
	myFunc := func(c clock.Clock) string {
		return c.Now().Format(time.RFC3339)
	}

	now := time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC)
	later := now.Add(time.Hour)
	// push 2 items to the queue
	// any extra items will block until consumed
	q.Push(now, later)

	// myFunc drains the time queue
	fmt.Println(myFunc(c))
	fmt.Println(myFunc(c))

	tc := make(chan string, 1)
	go func() {
		// and blocks when the queue is empty
		tc <- myFunc(c)
	}()

	select {
	case <-tc:
		// we shouldn't be seeing any values here because the queue is empty
		fmt.Println("oh no!")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("phew!")
	}

	evenlater := later.Add(time.Hour)
	// push another items to the queue
	q.Push(evenlater)

	t := <-tc
	fmt.Println(t)

	// Output:
	// 2018-12-31T00:00:00Z
	// 2018-12-31T01:00:00Z
	// phew!
	// 2018-12-31T02:00:00Z
}

func ExampleClock_NewTicker() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a test ticker (this ticker is returned to every clock.NewTicker())
	tt := clocktest.NewTicker()
	// ticker id to use when we have more than 1 ticker on the clock
	tid := 0 // can be anything
	// and a test clock that uses the ticker with some id
	c := clocktest.New(clocktest.WithTicker(tid, tt))

	// create a counter that increments every 1 second
	// and prints the received time
	ctr := &counter{
		c:   c,
		tid: tid,
		// use this to get notifications when the counter got incremented.
		// we can also use a time.Sleep() after a tt.Next(), but it's not guaranteed we get the timing right
		updated: make(chan struct{}, 1),
	}
	// run the counter in a goroutine
	go ctr.start(ctx)

	fmt.Println(ctr.Get())

	// advance the ticker
	now := time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC)
	tt.Next(ctx, now)

	// wait for update
	<-ctr.updated

	fmt.Println(ctr.Get())

	// advance the ticker again
	tt.Next(ctx, now.Add(time.Hour))

	// wait for update
	<-ctr.updated

	fmt.Println(ctr.Get())

	// Output:
	// 0
	// 2018-12-31T00:00:00Z
	// 1
	// 2018-12-31T01:00:00Z
	// 2
}

func ExampleClock_Sleep() {
	// create a sleeper
	s := clocktest.NewSleeper()
	// and a test clock that uses the sleeper
	c := clocktest.New(clocktest.WithSleeper(s))

	// trigger wakeup after 100ms-ish
	time.AfterFunc(100*time.Millisecond, func() {
		s.WakeUp(context.Background())
	})

	// sleep for 10s
	now := time.Now()
	c.Sleep(10 * time.Second)
	elapsed := time.Since(now)

	// check that we didn't actually sleep for 10s (at most 1s if OS is slow)
	if elapsed < time.Second {
		fmt.Println("hooray!")
	} else {
		fmt.Println("oops :/")
	}

	// Output:
	// hooray!
}

type counter struct {
	c   clock.Clock
	tid int

	mux sync.RWMutex
	v   int

	updated chan struct{}
}

func (c *counter) Get() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.v
}

func (c *counter) start(ctx context.Context) {
	ticker := c.c.NewTicker(time.Second, clock.TickerWithID(c.tid))
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ticker.C():
			c.mux.Lock()
			fmt.Println(t.Format(time.RFC3339Nano))
			c.v += 1
			c.mux.Unlock()
			c.updated <- struct{}{}
		}
	}
}
