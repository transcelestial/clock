package clock_test

import (
	"context"
	"fmt"
	"time"

	"github.com/transcelestial/clock"
)

func ExampleClock_Now() {
	c := clock.New()

	// inject the Clock interface as dependency into some func
	myFunc := func(c clock.Clock) {
		// use the clock the same way you'd use the "time" package
		fmt.Println(c.Now())
	}

	myFunc(c)
}

func ExampleClock_Sleep() {
	c := clock.New()

	now := time.Now()
	c.Sleep(time.Second)
	d := time.Since(now)

	if d >= time.Second {
		fmt.Println("slept a second")
	}

	// Output:
	// slept a second
}

func ExampleClock_NewTicker() {
	c := clock.New()

	// create a new ticker that ticks every 100ms
	ticker := c.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()

	var i int
	for {
		select {
		case <-ctx.Done():
			return
			// get notifications from the ticker (same way the "time" ticker.C chan works)
		case <-ticker.C():
			fmt.Println(i)
			i++
		}
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleClock_NewTimer() {
	c := clock.New()

	// create a new timer that will trigger after 100ms
	timer := c.NewTimer(100 * time.Millisecond)
	defer timer.Stop()

	// wait for the timer to trigger
	<-timer.C()

	fmt.Println("done")

	// Output:
	// done
}
