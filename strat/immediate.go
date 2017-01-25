package strat

import "time"

// ImmediateStrategy represents a scheduling strategy which will
// immediately execute a task.
type ImmediateStrategy struct {
}

// Immediate creates a new instance of the ImmediateStrategy which will
// activate at a specific time.
func Immediate() *ImmediateStrategy {
	return &ImmediateStrategy{}
}

// String returns the details of this strategy for human consumption.
func (s *ImmediateStrategy) String() string {
	return "immediate"
}

// Schedule returns a schedule which will result in the task being executed
// immediately.
func (s *ImmediateStrategy) Schedule() Schedule {
	return newImmediateSchedule()
}

type immediateSchedule struct {
	c chan time.Time
}

func newImmediateSchedule() *immediateSchedule {
	c := make(chan time.Time, 1)
	c <- time.Now()
	close(c)

	return &immediateSchedule{
		c: c,
	}
}

func (s *immediateSchedule) Events() <-chan time.Time {
	return s.c
}

func (s *immediateSchedule) Cancel() {
	c := make(chan time.Time)
	close(c)
	s.c = c
}
