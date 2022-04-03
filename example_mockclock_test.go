package clock_test

import (
	"sync"
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
	ck := mockclock.NewMockClock(ctrl)

	now := time.Now()
	ck.EXPECT().
		Now().
		Return(now)

	getTime := func(ck clock.Clock) time.Time {
		return ck.Now()
	}

	assert.Equal(t, now, getTime(ck))
}

func TestMockclockSleep(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)

	wait := make(chan struct{})
	ck.EXPECT().
		Sleep(gomock.Eq(time.Second)).
		Do(func(time.Duration) {
			<-wait
		})

	go func() {
		wait <- struct{}{}
	}()

	ck.Sleep(time.Second)
}

func TestMockclockNewTicker(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)
	mockticker := mockclock.NewMockTicker(ctrl)

	ck.EXPECT().
		NewTicker(gomock.Eq(time.Second)).
		Return(mockticker)

	next := make(chan time.Time)
	mockticker.EXPECT().
		Stop().
		After(mockticker.EXPECT().
			C().
			Return(next))

	ticker := ck.NewTicker(time.Second)
	defer ticker.Stop()

	now := time.Now()
	go func() {
		next <- now
	}()

	c := <-ticker.C()

	assert.Equal(t, now, c)
}

func TestMockclockNewTimer(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)
	mocktimer := mockclock.NewMockTimer(ctrl)

	ck.EXPECT().
		NewTimer(gomock.Eq(time.Second)).
		Return(mocktimer)

	next := make(chan time.Time)
	mocktimer.EXPECT().
		Stop().
		After(mocktimer.EXPECT().
			C().
			Return(next))

	timer := ck.NewTimer(time.Second)
	defer timer.Stop()

	now := time.Now()
	go func() {
		next <- now
	}()

	c := <-timer.C()

	assert.Equal(t, now, c)
}

func TestMockclockAfter(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)

	next := make(chan time.Time)
	ck.EXPECT().
		After(gomock.Eq(time.Second)).
		Return(next)

	now := time.Now()
	go func() {
		next <- now
	}()

	c := <-ck.After(time.Second)

	assert.Equal(t, now, c)
}

func TestMockclockAfterFunc(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)
	mocktimer := mockclock.NewMockTimer(ctrl)

	var wg sync.WaitGroup
	wg.Add(1)
	f := func() {
		wg.Done()
	}

	ck.EXPECT().
		AfterFunc(gomock.Eq(time.Second), gomock.Any()).
		Do(func(d time.Duration, f func()) {
			f()
		}).
		Return(mocktimer)

	_ = ck.AfterFunc(time.Second, f)

	wg.Wait()
}

func TestMockclockTick(t *testing.T) {
	defer leaktest.Check(t)()

	ctrl := gomock.NewController(t)
	ck := mockclock.NewMockClock(ctrl)

	next := make(chan time.Time)
	ck.EXPECT().
		Tick(gomock.Eq(time.Second)).
		Return(next)

	now := time.Now()
	go func() {
		next <- now
	}()

	c := <-ck.Tick(time.Second)

	assert.Equal(t, now, c)
}
