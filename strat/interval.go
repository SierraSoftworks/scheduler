package strat

import (
	"fmt"
	"time"
)

// IntervalStrategy is a scheduling strategy which triggers task
// execution at consistent intervals.
type IntervalStrategy struct {
	d time.Duration
}

// Every creates a new instance of the InvervalStrategy configured
// to emit events at regular intervals.
func Every(d time.Duration) *IntervalStrategy {
	return &IntervalStrategy{
		d: d,
	}
}

// String returns the details of this strategy for human consumption.
func (s *IntervalStrategy) String() string {
	return fmt.Sprintf("every %s", s.d.String())
}

// Interval returns the amount of time between consecutive executions.
func (s *IntervalStrategy) Interval() time.Duration {
	return s.d
}

// Schedule returns a schedule which will result in tasks being executed
// at fixed intervals.
func (s *IntervalStrategy) Schedule() Schedule {
	return newIntervalSchedule(s.d)
}

type intervalSchedule struct {
	c      chan time.Time
	cancel chan struct{}
}

func newIntervalSchedule(interval time.Duration) *intervalSchedule {
	c := make(chan time.Time)
	cancel := make(chan struct{})

	go func() {
		for run := true; run; {
			select {
			case <-cancel:
				run = false
			case t := <-time.After(interval):
				select {
				case c <- t:
				default:
				}
			}
		}

		close(c)
	}()

	return &intervalSchedule{
		c:      c,
		cancel: cancel,
	}
}

func (s *intervalSchedule) Events() <-chan time.Time {
	return s.c
}

func (s *intervalSchedule) Cancel() {
	select {
	case s.cancel <- struct{}{}:
	default:
	}
}
