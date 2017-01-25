package scheduler

import (
	"time"

	"github.com/SierraSoftworks/scheduler/strat"
)

// At instructs this task to execute at a specific time
// unless cancelled.
func (t *Task) At(time time.Time) *Task {
	return t.WithStrategy(strat.At(time))
}

// After instructs this task to execute after a specific
// delay unless cancelled.
func (t *Task) After(delay time.Duration) *Task {
	return t.WithStrategy(strat.Delay(delay))
}

// Every instructs this task to execute repeatedly at the given
// intervals until cancelled.
func (t *Task) Every(interval time.Duration) *Task {
	return t.WithStrategy(strat.Every(interval))
}

// Never instructs this task to never execute, this is the default
// mode of operation.
func (t *Task) Never() *Task {
	return t.WithStrategy(strat.Never())
}

// Manual will instruct this task to execute when the provided
// manual strategy is triggered.
func (t *Task) Manual(s *strat.ManualStrategy) *Task {
	return t.WithStrategy(s)
}

// Immediately will instruct this task to execute immediately when
// scheduled.
func (t *Task) Immediately() *Task {
	return t.WithStrategy(strat.Immediate())
}
