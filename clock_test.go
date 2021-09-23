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

func TestSysClockTimer(t *testing.T) {
	c := New()
	timer := c.NewTimer(time.Millisecond)
	if assert.NotNil(t, timer) {
		timer.Stop()
	}
}

func TestSysClockSleep(t *testing.T) {
	c := New()
	before := time.Now()
	c.Sleep(100 * time.Millisecond)
	elapsed := time.Since(before)
	assert.True(t, elapsed >= 100*time.Millisecond)
}
