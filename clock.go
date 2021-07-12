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
	// Note that the options are needed when testing, to set the ticker ID, and it won't have any effect when using the
	// system clock.
	NewTicker(d time.Duration, opts ...TickerOption) Ticker
	// Sleep pauses execution in the current thread for d duration.
	Sleep(d time.Duration)
}

type TickerOption func(*TickerOptions)

type TickerOptions struct {
	ID interface{}
}

func TickerWithID(id interface{}) TickerOption {
	return func(o *TickerOptions) {
		o.ID = id
	}
}

type sysClock struct{}

func (c *sysClock) Now() time.Time {
	return time.Now()
}

func (c *sysClock) NewTicker(d time.Duration, opts ...TickerOption) Ticker {
	return newSysTicker(d)
}

func (c *sysClock) Sleep(d time.Duration) {
	time.Sleep(d)
}
