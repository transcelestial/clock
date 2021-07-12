package clocktest

import (
	"time"
)

// NewTimeQueue creates a time.Time queue of a given size
func NewTimeQueue(size int, times ...time.Time) *TimeQueue {
	return &TimeQueue{time: make(chan time.Time, size)}
}

type TimeQueue struct {
	time chan time.Time
}

// Drain removes the first time.Time from the queue.
// This call will block if there's no items in the queue.
func (q *TimeQueue) Drain() time.Time {
	return <-q.time
}

// Push pushes new time.Time items to the queue.
// This call blocks when the queue is full.
func (q *TimeQueue) Push(times ...time.Time) {
	for _, t := range times {
		q.time <- t
	}
}
