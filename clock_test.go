package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSysClockNow(t *testing.T) {
	c := New()
	now := c.Now()
	if assert.NotNil(t, now) {
		assert.True(t, time.Now().After(now))
	}
}

func TestSysClockTicker(t *testing.T) {
	c := New()
	ticker := c.NewTicker(time.Millisecond)
	if assert.NotNil(t, ticker) {
		ticker.Stop()
	}
}

func TestSysClockSleep(t *testing.T) {
	c := New()
	before := time.Now()
	c.Sleep(100 * time.Millisecond)
	elapsed := time.Since(before)
	assert.True(t, elapsed >= 100*time.Millisecond)
}

func TestTickerOptions(t *testing.T) {
	opts := &TickerOptions{}
	TickerWithID(123)(opts)
	assert.Equal(t, 123, opts.ID)
}
