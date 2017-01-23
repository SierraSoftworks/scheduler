package scheduler

import (
	"fmt"
	"time"
)

// IntervalStrategy is a scheduling strategy which triggers task
// execution at consistent intervals.
type IntervalStrategy struct {
	d time.Duration
}

// String returns the details of this strategy for human consumption.
func (s *IntervalStrategy) String() string {
	return fmt.Sprintf("every %s", s.d.String())
}

// Interval returns the amount of time between consecutive executions.
func (s *IntervalStrategy) Interval() time.Duration {
	return s.d
}

// Next returns a channel which will emit a message the next time a
// task should be executed.
func (s *IntervalStrategy) Next() <-chan time.Time {
	return time.After(s.d)
}

// Every will configure this task to run at regular intervals
func (t *Task) Every(d time.Duration) *Task {
	return t.WithStrategy(&IntervalStrategy{
		d: d,
	})
}
