package strat

import "time"

// NeverStrategy represents a scheduling strategy which will
// never trigger the execution of a task.
type NeverStrategy struct {
}

// String returns the details of this strategy for human consumption.
func (s *NeverStrategy) String() string {
	return "never"
}

// Schedule returns a schedule which will result in no tasks ever being
// executed.
func (s *NeverStrategy) Schedule() Schedule {
	return newNeverSchedule()
}

// Never creates a new scheduling strategy which will never execute a task.
func Never() *NeverStrategy {
	return &NeverStrategy{}
}

type neverSchedule struct {
	c chan time.Time
}

func newNeverSchedule() *neverSchedule {
	c := make(chan time.Time)
	close(c)

	return &neverSchedule{
		c: c,
	}
}

func (s *neverSchedule) Events() <-chan time.Time {
	return s.c
}

func (s *neverSchedule) Cancel() {

}
