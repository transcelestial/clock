package clock

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSysClockNow(t *testing.T) {
	ck := New()
	now := ck.Now()
	require.NotNil(t, now)
	assert.True(t, time.Now().After(now))
}

func TestSysClockTicker(t *testing.T) {
	ck := New()
	d := osDelta()
	before := time.Now()
	ticker := ck.NewTicker(d)
	require.NotNil(t, ticker)
	c := ticker.C()
	require.NotNil(t, c)
	<-c
	require.True(t, time.Since(before) >= d)
	ticker.Reset(d)
	<-c
	assert.True(t, time.Since(before) >= 2*d)
	ticker.Stop()
}

func TestSysClockTimer(t *testing.T) {
	ck := New()
	d := osDelta()
	before := time.Now()
	timer := ck.NewTimer(d)
	require.NotNil(t, timer)
	c := timer.C()
	require.NotNil(t, c)
	<-c
	require.True(t, time.Since(before) >= d)
	timer.Reset(d)
	<-c
	assert.True(t, time.Since(before) >= 2*d)
	timer.Stop()
}

func TestSysClockSleep(t *testing.T) {
	ck := New()
	before := time.Now()
	d := osDelta()
	ck.Sleep(d)
	assert.True(t, time.Since(before) >= d)
}

func TestSysClockAfter(t *testing.T) {
	ck := New()
	d := osDelta()
	before := time.Now()
	c := ck.After(d)
	require.NotNil(t, c)
	<-c
	assert.True(t, time.Since(before) >= d)
}

func TestSysClockAfterFunc(t *testing.T) {
	ck := New()
	d := osDelta()
	before := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		wg.Done()
	}
	timer := ck.AfterFunc(d, f)
	require.NotNil(t, timer)
	c := timer.C()
	require.Nil(t, c)
	timer.Reset(d * 2)
	wg.Wait()
	assert.True(t, time.Since(before) >= d*2)
	timer.Stop()
}

func TestSysClockTick(t *testing.T) {
	ck := New()
	d := osDelta()
	before := time.Now()
	c := ck.Tick(d)
	require.NotNil(t, c)
	<-c
	assert.True(t, time.Since(before) >= d)
}
