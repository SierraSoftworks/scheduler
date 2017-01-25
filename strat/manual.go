package strat

import (
	"sync"
	"time"
)

// ManualStrategy provides the ability to manually trigger tasks by
// calling the Trigger method on this strategy.
type ManualStrategy struct {
	targets []chan<- time.Time
	l       sync.Mutex
}

// Manual creates a new instance of the ManualStrategy, allowing you
// to manually trigger the execution of tasks by calling Trigger().
func Manual() *ManualStrategy {
	return &ManualStrategy{
		targets: []chan<- time.Time{},
	}
}

// String returns the details of this strategy for human consumption.
func (s *ManualStrategy) String() string {
	return "manual"
}

// Schedule returns a schedule which will result in tasks being executed
// when Trigger() is called.
func (s *ManualStrategy) Schedule() Schedule {
	return newManualSchedule(s)
}

// Trigger will cause all dependent tasks to be executed whenever it is
// called.
func (s *ManualStrategy) Trigger() {
	s.l.Lock()
	defer s.l.Unlock()

	t := time.Now()
	for _, target := range s.targets {
		select {
		case target <- t:
		default:
		}
	}
}

func (s *ManualStrategy) addTarget(c chan<- time.Time) {
	s.l.Lock()
	defer s.l.Unlock()

	s.targets = append(s.targets, c)
}

func (s *ManualStrategy) removeTarget(c chan<- time.Time) {
	s.l.Lock()
	defer s.l.Unlock()

	for i, t := range s.targets {
		if t == c {
			s.targets[i] = s.targets[len(s.targets)-1]
			s.targets = s.targets[:len(s.targets)-1]
			return
		}
	}
}

type manualSchedule struct {
	c chan time.Time

	s *ManualStrategy
}

func newManualSchedule(s *ManualStrategy) *manualSchedule {
	c := make(chan time.Time)
	s.addTarget(c)
	return &manualSchedule{
		c: c,
		s: s,
	}
}

func (s *manualSchedule) Events() <-chan time.Time {
	return s.c
}

func (s *manualSchedule) Cancel() {
	s.s.removeTarget(s.c)
	close(s.c)
}
