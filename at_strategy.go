package scheduler

import (
	"fmt"
	"time"
)

// AtStrategy is a scheduling strategy which triggers task
// execution at consistent intervals.
type AtStrategy struct {
	t time.Time
}

// String returns the details of this strategy for human consumption.
func (s *AtStrategy) String() string {
	return fmt.Sprintf("at %s", s.t.String())
}

// Time returns the time at which the operation is scheduled to take place.
func (s *AtStrategy) Time() time.Time {
	return s.t
}

// Next returns a channel which will emit a message the next time a
// task should be executed.
func (s *AtStrategy) Next() <-chan time.Time {
	d := s.t.Sub(time.Now())
	if d < 0 {
		c := make(chan time.Time)
		close(c)
		return c
	}

	return time.After(d)
}

// At creates a new scheduling strategy which will trigger
// the execution of tasks at the specified time.
func (t *Task) At(time time.Time) *Task {
	return t.WithStrategy(&AtStrategy{
		t: time,
	})
}
