package scheduler

import "time"

// NeverStrategy represents a scheduling strategy which will
// never trigger the execution of a task.
type NeverStrategy struct {
}

// String returns the details of this strategy for human consumption.
func (s *NeverStrategy) String() string {
	return "never"
}

// Next returns a channel which will emit a message the next time a
// task should be executed.
func (s *NeverStrategy) Next() <-chan time.Time {
	c := make(chan time.Time)
	close(c)
	return c
}

// Never will instruct this task to never trigger.
func (t *Task) Never() *Task {
	return t.WithStrategy(&NeverStrategy{})
}
