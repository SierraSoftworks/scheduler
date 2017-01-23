package scheduler

import (
	"fmt"
	"time"
)

// ActiveTask represents a task which has been scheduled
// for execution according to a schedule.
type ActiveTask struct {
	*Task

	cancel chan struct{}

	errors    chan error
	lastError error
}

func newActiveTask(task *Task) *ActiveTask {
	return &ActiveTask{
		Task: task,

		cancel: make(chan struct{}),

		errors:    make(chan error),
		lastError: nil,
	}
}

// Cancel will abort any further executions of this specific
// task.
func (t *ActiveTask) Cancel() {
	select {
	case t.cancel <- struct{}{}:
	default:
	}
}

// CancelWhen will abort any further executions of this specific
// task after the provided channel emits a message.
func (t *ActiveTask) CancelWhen(c <-chan time.Time) {
	go func() {
		_, ok := <-c
		if !ok {
			return
		}

		t.Cancel()
	}()
}

// LastError returns the last error encountered when executing
// this task.
func (t *ActiveTask) LastError() error {
	return t.lastError
}

// Errors returns the error channel for this task, on which all
// execution errors will be emitted.
func (t *ActiveTask) Errors() <-chan error {
	return t.errors
}

// String returns a string representation of this active task
func (t *ActiveTask) String() string {
	return fmt.Sprintf("Active %s", t.Task.String())
}

func (t *ActiveTask) run() {
	go func() {
		for {
			select {
			case <-t.cancel:
				return
			case ts, ok := <-t.strategy.Next():
				if !ok {
					return
				}

				if err := t.action(ts); err != nil {
					t.lastError = err
					select {
					case t.errors <- err:
					default:
					}
				}
			}
		}
	}()
}
