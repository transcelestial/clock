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
}

type sysClock struct{}

func (c *sysClock) Now() time.Time {
	return time.Now()
}

func (c *sysClock) NewTicker(d time.Duration) Ticker {
	return newSysTicker(d)
}
