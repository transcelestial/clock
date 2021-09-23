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
	NewTicker(d time.Duration) Ticker
	// NewTimer creates a new timer that uses the system clock (same as `time.NewTimer()`).
	NewTimer(d time.Duration) Timer
	// Sleep pauses execution in the current thread for d duration (same as `time.Sleep()`).
	Sleep(d time.Duration)
}

type sysClock struct{}

func (c *sysClock) Now() time.Time {
	return time.Now()
}

func (c *sysClock) NewTicker(d time.Duration) Ticker {
	return newSysTicker(d)
}

func (c *sysClock) NewTimer(d time.Duration) Timer {
	return newSysTimer(d)
}

func (c *sysClock) Sleep(d time.Duration) {
	time.Sleep(d)
}
