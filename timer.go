package clock

import "time"

func newSysTimer(d time.Duration, f func()) Timer {
	var t *time.Timer
	if f != nil {
		t = time.AfterFunc(d, f)
	} else {
		t = time.NewTimer(d)
	}
	return &sysTimer{t}
}

type Timer interface {
	// C returns the timer.C chan.
	C() <-chan time.Time
	// Reset resets the timer to a different duration.
	Reset(d time.Duration)
	// Stop stops the timer.
	Stop() bool
}

type sysTimer struct {
	t *time.Timer
}

func (t *sysTimer) C() <-chan time.Time {
	return t.t.C
}

func (t *sysTimer) Stop() bool {
	return t.t.Stop()
}

func (t *sysTimer) Reset(d time.Duration) {
	t.t.Reset(d)
}
