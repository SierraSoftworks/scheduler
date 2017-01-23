package scheduler

import (
	"fmt"
	"time"
)

// Task represents a task definition which can be scheduled
// by calling the Schedule method.
type Task struct {
	action   func(t time.Time) error
	strategy Strategy
}

// Do creates a new task definition which can be customized
// before calling Schedule to begin executing the task according
// to the specified schedule.
func Do(action func(t time.Time) error) *Task {
	return &Task{
		action:   action,
		strategy: &NeverStrategy{},
	}
}

// WithStrategy allows you tp specify the strategy to be used
// for scheduling execution of this task.
func (t *Task) WithStrategy(s Strategy) *Task {
	t.strategy = s
	return t
}

// Schedule will create a scheduled task from this task
// definition and start scheduling operations according
// to the strategy defined in this task.
func (t *Task) Schedule() *ActiveTask {
	s := newActiveTask(t)
	s.run()
	return s
}

// String returns a human readable representation of this
// task.
func (t *Task) String() string {
	return fmt.Sprintf("Task(%s)", t.strategy.String())
}
