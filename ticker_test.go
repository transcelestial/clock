package clock

import (
	"testing"
	"time"
)

// Do a sanity check
// https://golang.org/src/time/tick_test.go
func TestSysTicker(t *testing.T) {
	d := osDelta()
	ticker := newSysTicker(d)
	<-ticker.C()
	ticker.Stop()
	time.Sleep(2 * d)
	select {
	case <-ticker.C():
		t.FailNow()
	default:
		// ok
	}
}
