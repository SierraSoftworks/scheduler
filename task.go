package scheduler

import (
	"fmt"
	"time"

	"github.com/SierraSoftworks/scheduler/strat"
)

// Task represents a task definition which can be scheduled
// by calling the Schedule method.
type Task struct {
	action   func(t time.Time) error
	strategy strat.Strategy

	errorsHandler func(c <-chan error)
}

// New creates a new task definition which can then be customized
// to provide any behavior you wish. It is more idiomatic to use the
// Do() function to prime a task than to construct it manually.
func New() *Task {
	return &Task{
		action: func(t time.Time) error {
			return nil
		},
		strategy:      strat.Never(),
		errorsHandler: func(c <-chan error) {},
	}
}

// Do creates a new task definition which can be customized
// before calling Schedule to begin executing the task according
// to the specified schedule.
func Do(action func(t time.Time) error) *Task {
	return New().Do(action)
}

// Strategy returns the currently configured strategy for this
// task.
func (t *Task) Strategy() strat.Strategy {
	return t.strategy
}

// Action returns the currently configured action for this task.
func (t *Task) Action() func(time.Time) error {
	return t.action
}

// Handler returns the error handler currently configured for
// this task.
func (t *Task) Handler() func(<-chan error) {
	return t.errorsHandler
}

// Clone returns a shallow clone of the current task.
func (t *Task) Clone() *Task {
	return &Task{
		action:        t.action,
		strategy:      t.strategy,
		errorsHandler: t.errorsHandler,
	}
}

// Do allows you to change the function that this task will
// execute. It will return a clone of the current task with
// the provided action.
func (t *Task) Do(action func(t time.Time) error) *Task {
	if action == nil {
		action = func(ts time.Time) error {
			return nil
		}
	}

	t2 := t.Clone()
	t2.action = action
	return t2
}

// WithStrategy allows you to specify the strategy to be used
// for scheduling execution of this task. It will return a
// clone of the current task with the provided strategy.
func (t *Task) WithStrategy(s strat.Strategy) *Task {
	if s == nil {
		s = strat.Never()
	}

	t2 := t.Clone()
	t2.strategy = s
	return t2
}

// WithHandler allows you to specify the error handler to be
// used for tracking errors which occur during execution of
// this task.
func (t *Task) WithHandler(h func(<-chan error)) *Task {
	if h == nil {
		h = func(c <-chan error) {}
	}

	t2 := t.Clone()
	t2.errorsHandler = h
	return t2
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
