package clock

import (
	"runtime"
	"testing"
	"time"
)

// Do a sanity check
func TestSysTimer(t *testing.T) {
	delta := 20 * time.Millisecond

	if (runtime.GOOS == "darwin" || runtime.GOOS == "ios") && runtime.GOARCH == "arm64" {
		delta = 100 * time.Millisecond
	}

	ticker := newSysTimer(delta)
	<-ticker.C()
	ticker.Stop()
	time.Sleep(2 * delta)
	select {
	case <-ticker.C():
		t.FailNow()
	default:
		// ok
	}
}
