package clock

import "time"

func newSysTimer(d time.Duration) Timer {
	t := time.NewTimer(d)
	return &sysTimer{t}
}

type Timer interface {
	C() <-chan time.Time
	Reset(d time.Duration)
	Stop()
}

type sysTimer struct {
	t *time.Timer
}

func (t *sysTimer) C() <-chan time.Time {
	return t.t.C
}

func (t *sysTimer) Stop() {
	t.t.Stop()
}

func (t *sysTimer) Reset(d time.Duration) {
	t.t.Reset(d)
}
