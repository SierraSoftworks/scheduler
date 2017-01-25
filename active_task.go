package scheduler

import (
	"fmt"
	"time"

	"sync"

	"github.com/SierraSoftworks/scheduler/strat"
)

// ActiveTask represents a task which has been scheduled
// for execution according to a schedule.
type ActiveTask struct {
	task *Task

	schedule  strat.Schedule
	errors    chan error
	lastError error
	running   sync.WaitGroup
}

func newActiveTask(task *Task) *ActiveTask {
	return &ActiveTask{
		task: task,

		schedule:  task.strategy.Schedule(),
		errors:    make(chan error),
		lastError: nil,
	}
}

// Cancel will abort any further executions of this specific
// task.
func (t *ActiveTask) Cancel() {
	t.schedule.Cancel()
}

// CancelWhen will abort any further executions of this specific
// task after the provided channel emits a message.
func (t *ActiveTask) CancelWhen(c <-chan time.Time) *ActiveTask {
	if c == nil {
		return t
	}

	go func() {
		_, ok := <-c
		if !ok {
			return
		}

		t.Cancel()
	}()

	return t
}

// LastError returns the last error encountered when executing
// this task.
func (t *ActiveTask) LastError() error {
	return t.lastError
}

// String returns a string representation of this active task
func (t *ActiveTask) String() string {
	return fmt.Sprintf("Active %s", t.task.String())
}

// Wait will block until this scheduled task has completed execution.
// In the case of infinitely repeating tasks, this will block until
// Cancel is called.
func (t *ActiveTask) Wait() {
	t.running.Wait()
}

// Done returns a channel which you can use to determine when this task
// has completed execution.
func (t *ActiveTask) Done() <-chan time.Time {
	c := make(chan time.Time)

	go func() {
		t.running.Wait()
		c <- time.Now()
		close(c)
	}()

	return c
}

func (t *ActiveTask) run() {
	t.running.Add(1)

	// Schedule the error handler on a new goroutine
	// This enables you to write very simple error handlers (a for loop)
	// without breaking things. The downside is that we need to prevent
	// errors from being emitted before this goroutine is started.
	// In practice that's not much of a problem.
	var errorHandlerReady sync.WaitGroup
	errorHandlerReady.Add(1)
	go func() {
		errorHandlerReady.Done()
		t.task.errorsHandler(t.errors)
	}()

	go func() {

		for ts := range t.schedule.Events() {
			t.running.Add(1)
			go func(ts time.Time) {
				err := t.task.action(ts)
				if err != nil {
					t.lastError = err

					errorHandlerReady.Wait()
					select {
					case t.errors <- err:
					default:
					}
				}

				t.running.Done()
			}(ts)
		}

		t.running.Done()

		t.running.Wait()
		close(t.errors)
	}()
}
