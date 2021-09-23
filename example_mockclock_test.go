package clock_test

import (
	"testing"
	"time"

	"github.com/fortytw2/leaktest"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/transcelestial/clock"
	"github.com/transcelestial/clock/mockclock"
)

func TestMockclockNow(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mockclock.NewMockClock(ctrl)

	now := time.Now()
	c.EXPECT().
		Now().
		Return(now)

	getTime := func(c clock.Clock) time.Time {
		return c.Now()
	}

	assert.Equal(t, now, getTime(c))
}

func TestMockclockSleep(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	c := mockclock.NewMockClock(ctrl)

	wait := make(chan struct{})
	c.EXPECT().
		Sleep(gomock.Eq(time.Second)).
		Do(func(time.Duration) {
			<-wait
		})

	go func() {
		wait <- struct{}{}
	}()

	c.Sleep(time.Second)
}

func TestMockclockTicker(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	c := mockclock.NewMockClock(ctrl)
	mockticker := mockclock.NewMockTicker(ctrl)

	c.EXPECT().
		NewTicker(gomock.Eq(time.Second)).
		Return(mockticker)

	next := make(chan time.Time)
	mockticker.EXPECT().
		Stop().
		After(mockticker.EXPECT().
			C().
			Return(next))

	ticker := c.NewTicker(time.Second)
	defer ticker.Stop()

	now := time.Now()
	go func() {
		next <- now
	}()

	tm := <-ticker.C()

	assert.Equal(t, now, tm)
}

func TestMockclockTimer(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	c := mockclock.NewMockClock(ctrl)
	mocktimer := mockclock.NewMockTimer(ctrl)

	c.EXPECT().
		NewTimer(gomock.Eq(time.Second)).
		Return(mocktimer)

	next := make(chan time.Time)
	mocktimer.EXPECT().
		Stop().
		After(mocktimer.EXPECT().
			C().
			Return(next))

	timer := c.NewTimer(time.Second)
	defer timer.Stop()

	now := time.Now()
	go func() {
		next <- now
	}()

	tm := <-timer.C()

	assert.Equal(t, now, tm)
}
