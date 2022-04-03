package clock

import "time"

func newSysTicker(d time.Duration) Ticker {
	t := time.NewTicker(d)
	return &sysTicker{t}
}

type Ticker interface {
	// C returns the ticker.C chan.
	C() <-chan time.Time
	// Reset resets the ticker to a different duration.
	Reset(d time.Duration)
	// Stop stops the ticker.
	Stop()
}

type sysTicker struct {
	t *time.Ticker
}

func (t *sysTicker) C() <-chan time.Time {
	return t.t.C
}

func (t *sysTicker) Stop() {
	t.t.Stop()
}

func (t *sysTicker) Reset(d time.Duration) {
	t.t.Reset(d)
}
