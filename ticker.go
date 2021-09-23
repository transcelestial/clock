package clock

import "time"

func newSysTicker(d time.Duration) Ticker {
	t := time.NewTicker(d)
	return &sysTicker{t}
}

type Ticker interface {
	C() <-chan time.Time
	Reset(d time.Duration)
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
