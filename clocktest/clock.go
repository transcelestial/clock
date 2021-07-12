package clocktest

import (
	"time"

	"github.com/transcelestial/clock"
)

// New creates a new test clock.
// Depending on your use case, you must provide at least one option (a ticker, time queue, etc).
func New(options ...Option) clock.Clock {
	opts := &Options{}
	for _, o := range options {
		o(opts)
	}
	return &testClock{
		t:  opts.ticker,
		tq: opts.tq,
		s:  opts.s,
	}
}

type Options struct {
	ticker clock.Ticker
	tq     *TimeQueue
	s      *Sleeper
}

type Option func(*Options)

// WithTicker sets a ticker on the clock.
func WithTicker(t clock.Ticker) Option {
	return func(o *Options) {
		o.ticker = t
	}
}

// WithTimeQueue sets a time queue on the clock.
func WithTimeQueue(tq *TimeQueue) Option {
	return func(o *Options) {
		o.tq = tq
	}
}

// WithSleeper sets a sleeper on the clock.
func WithSleeper(s *Sleeper) Option {
	return func(o *Options) {
		o.s = s
	}
}

type testClock struct {
	t  clock.Ticker
	tq *TimeQueue
	s  *Sleeper
}

// Now returns the first time.Time from the time queue
func (c *testClock) Now() time.Time {
	return c.tq.Drain()
}

// NewTicker returns the ticker the clock was initialized with
func (c *testClock) NewTicker(d time.Duration) clock.Ticker {
	return c.t
}

func (c *testClock) Sleep(d time.Duration) {
	c.s.Sleep()
}

var _ clock.Clock = &testClock{}
