package clock

import (
	"runtime"
	"time"
)

func osDelta() time.Duration {
	delta := 20 * time.Millisecond

	if (runtime.GOOS == "darwin" || runtime.GOOS == "ios") && runtime.GOARCH == "arm64" {
		delta = 100 * time.Millisecond
	}

	return delta
}
