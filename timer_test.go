package clock

import (
	"sync"
	"testing"
	"time"
)

// Do a sanity check
func TestSysTimer(t *testing.T) {
	d := osDelta()
	timer := newSysTimer(d, nil)
	<-timer.C()
	timer.Stop()
	time.Sleep(2 * d)
	select {
	case <-timer.C():
		t.FailNow()
	default:
		// ok
	}
}

func TestSysTimerCb(t *testing.T) {
	d := osDelta()
	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		wg.Done()
	}
	timer := newSysTimer(d, f)
	defer timer.Stop()
	wg.Wait()
}
