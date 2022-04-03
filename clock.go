package clock

import "time"

// New creates a new system clock
func New() Clock {
	return &sysClock{}
}

type Clock interface {
	// Now returns the current system time (same as `time.Now()`)
	Now() time.Time
	// NewTicker creates a new ticker that uses the system clock (same as `time.NewTicker()`).
	NewTicker(time.Duration) Ticker
	// NewTimer creates a new timer that uses the system clock (same as `time.NewTimer()`).
	NewTimer(time.Duration) Timer
	// Sleep pauses execution in the current thread for d duration (same as `time.Sleep()`).
	Sleep(time.Duration)
	// After waits for the duration to elapse and then sends the current time
	// on the returned channel (same as `time.After()`).
	After(time.Duration) <-chan time.Time
	// AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine (same as `time.AfterFunc()`).
	AfterFunc(time.Duration, func()) Timer
	// Tick is a convenience wrapper for NewTicker providing access to the ticking
	// channel only (same as `time.Tick()`).
	Tick(time.Duration) <-chan time.Time
}

type sysClock struct{}

func (ck *sysClock) Now() time.Time {
	return time.Now()
}

func (ck *sysClock) NewTicker(d time.Duration) Ticker {
	return newSysTicker(d)
}

func (ck *sysClock) NewTimer(d time.Duration) Timer {
	return newSysTimer(d, nil)
}

func (ck *sysClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (ck *sysClock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (ck *sysClock) AfterFunc(d time.Duration, f func()) Timer {
	return newSysTimer(d, f)
}

func (ck *sysClock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
