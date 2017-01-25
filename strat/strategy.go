package strat

// Strategy represents a scheduling strategy used to trigger
// the execution of a task.
type Strategy interface {
	// String gives a short description of this strategy.
	String() string

	// Schedule retrieves a new schedule instance which can be
	// used by an executor to run tasks according to this strategy's
	// requirements.
	Schedule() Schedule
}
