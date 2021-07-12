package clocktest

import (
	"fmt"
	"sync"
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
	return &testClock{opts: opts}
}

type Options struct {
	tickers sync.Map
	tq      *TimeQueue
	s       *Sleeper
}

type Option func(*Options)

// WithTicker sets a ticker on the clock.
func WithTicker(id interface{}, t clock.Ticker) Option {
	return func(o *Options) {
		if _, ok := o.tickers.Load(id); ok {
			panic(fmt.Sprintf("ticker with id %v already exists", id))
		}
		o.tickers.Store(id, t)
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
	opts *Options
}

// Now returns the first time.Time from the time queue
func (c *testClock) Now() time.Time {
	return c.opts.tq.Drain()
}

// NewTicker returns the ticker for the given id
func (c *testClock) NewTicker(d time.Duration, opts ...clock.TickerOption) clock.Ticker {
	tkopts := &clock.TickerOptions{}
	for _, o := range opts {
		o(tkopts)
	}
	v, ok := c.opts.tickers.Load(tkopts.ID)
	if !ok {
		panic(fmt.Sprintf("no ticker found for id %v: ", tkopts.ID))
	}
	return v.(clock.Ticker)
}

func (c *testClock) Sleep(d time.Duration) {
	c.opts.s.Sleep()
}

var _ clock.Clock = &testClock{}
