package strat

import "time"

// Schedule is the interface that a scheduler provides to enable
// consumers to perform executions on specific events, or stop the
// schedule at any time.
type Schedule interface {
	// Events returns a channel which emits a message whenever a
	// task is scheduled to be performed. It includes the current
	// timestamp in that message to provide context to an execution.
	Events() <-chan time.Time

	// Cancel should prevent any subsequent messages from being posted
	// to the Events channel and close the channel.
	Cancel()
}
