package strat

import (
	"fmt"
	"time"
)

// AtStrategy is a scheduling strategy which triggers task
// execution at a specific time.
type AtStrategy struct {
	t time.Time
}

// At creates a new instance of an AtStrategy which will activate
// at a specific time.
func At(t time.Time) *AtStrategy {
	return &AtStrategy{
		t: t,
	}
}

// String returns the details of this strategy for human consumption.
func (s *AtStrategy) String() string {
	return fmt.Sprintf("at %s", s.t.String())
}

// Time returns the time at which the operation is scheduled to take place.
func (s *AtStrategy) Time() time.Time {
	return s.t
}

// Schedule returns a new schedule which will result in a task being executed
// at the specified time.
func (s *AtStrategy) Schedule() Schedule {
	d := s.t.Sub(time.Now())
	if d < 0 {
		return newNeverSchedule()
	}

	return newDelaySchedule(d)
}
