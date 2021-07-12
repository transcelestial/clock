package clocktest

import "context"

func NewSleeper() *Sleeper {
	return &Sleeper{c: make(chan struct{})}
}

type Sleeper struct {
	c chan struct{}
}

func (s *Sleeper) Sleep() {
	<-s.c
}

func (s *Sleeper) WakeUp(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	case s.c <- struct{}{}:
	}
}
