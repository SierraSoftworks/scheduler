package scheduler

import "time"

// Strategy represents a scheduling strategy used to trigger
// the execution of a task.
type Strategy interface {
	// String gives a short description of this strategy.
	String() string

	// Next returns a channel which emits the current time
	// whenever it wishes to trigger the execution of a task.
	// Tasks are triggered when the channel emits a value, the
	// time emitted is passed to the task action, allowing it
	// to convey data for algorithms which require the scheduled
	// timestamp.
	Next() <-chan time.Time
}
