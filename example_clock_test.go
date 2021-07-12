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
