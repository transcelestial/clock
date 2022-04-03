package clock_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/transcelestial/clock"
)

func ExampleClock_Now() {
	ck := clock.New()

	// inject the Clock interface as dependency into some func
	myFunc := func(ck clock.Clock) {
		// use the clock the same way you'd use the "time" package
		fmt.Println(ck.Now())
	}

	myFunc(ck)
}

func ExampleClock_Sleep() {
	ck := clock.New()

	now := time.Now()
	ck.Sleep(time.Second)
	d := time.Since(now)

	if d >= time.Second {
		fmt.Println("slept a second")
	}

	// Output:
	// slept a second
}

func ExampleClock_NewTicker() {
	ck := clock.New()

	// create a new ticker that ticks every 100ms
	ticker := ck.NewTicker(100 * time.Millisecond)
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
	ck := clock.New()

	// create a new timer that will trigger after 100ms
	timer := ck.NewTimer(100 * time.Millisecond)
	defer timer.Stop()

	// wait for the timer to trigger
	<-timer.C()

	fmt.Println("done")

	// Output:
	// done
}

func ExampleClock_After() {
	ck := clock.New()

	// wait for the timer to trigger
	<-ck.After(100 * time.Millisecond)

	fmt.Println("done")

	// Output:
	// done
}

func ExampleClock_AfterFunc() {
	ck := clock.New()

	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		wg.Done()
	}

	_ = ck.AfterFunc(100*time.Millisecond, f)

	// wait for the timer to trigger
	wg.Wait()

	fmt.Println("done")

	// Output:
	// done
}

func ExampleClock_Tick() {
	ck := clock.New()

	// wait for the timer to trigger
	<-ck.Tick(100 * time.Millisecond)

	fmt.Println("done")

	// Output:
	// done
}
