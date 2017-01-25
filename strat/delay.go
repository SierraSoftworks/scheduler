package strat

import (
	"fmt"
	"time"
)

// DelayStrategy is a scheduling strategy which triggers task
// execution after a specific amount of time.
type DelayStrategy struct {
	d time.Duration
}

// Delay creates a new instance of an AfterStrategy which will activate
// after a specific duration.
func Delay(d time.Duration) *DelayStrategy {
	return &DelayStrategy{
		d: d,
	}
}

// String returns the details of this strategy for human consumption.
func (s *DelayStrategy) String() string {
	return fmt.Sprintf("after %s", s.d.String())
}

// Time returns the time at which the operation is scheduled to take place.
func (s *DelayStrategy) Delay() time.Duration {
	return s.d
}

// Schedule returns a new schedule which will result in a task being executed
// after the specified amount of time.
func (s *DelayStrategy) Schedule() Schedule {
	if s.d < 0 {
		return newNeverSchedule()
	}

	return newDelaySchedule(s.d)
}

type delaySchedule struct {
	c      chan time.Time
	cancel chan struct{}
}

func newDelaySchedule(d time.Duration) *delaySchedule {
	c := make(chan time.Time)
	cancel := make(chan struct{})

	go func() {
		select {
		case <-cancel:
		case t := <-time.After(d):
			select {
			case c <- t:
			default:
			}
		}

		close(c)
	}()

	return &delaySchedule{
		c:      c,
		cancel: cancel,
	}
}

func (s *delaySchedule) Events() <-chan time.Time {
	return s.c
}

func (s *delaySchedule) Cancel() {
	select {
	case s.cancel <- struct{}{}:
	default:
	}
}
