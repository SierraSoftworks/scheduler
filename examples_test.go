package scheduler

import (
	"fmt"
	"math/rand"
	"time"
)

func ExampleActiveTask_CancelWhen() {
	t1 := Do(func(t time.Time) error {
		fmt.Println("Task 1 won the race")
		return nil
	}).
		After(time.Duration(5+rand.Intn(30)) * time.Millisecond).
		Schedule()

	t2 := Do(func(t time.Time) error {
		fmt.Println("Task 2 won the race")
		return nil
	}).
		After(time.Duration(5+rand.Intn(30)) * time.Millisecond).
		Schedule().
		CancelWhen(t1.Done())

	t1.CancelWhen(t2.Done())

	select {
	case <-t1.Done():
	case <-t2.Done():
	}

	fmt.Println("Race is over...")
}

func ExampleTask_After() {
	Do(func(t time.Time) error {
		fmt.Println("executing scheduled task")
		return nil
	}).
		After(5 * time.Second).
		Schedule().
		Wait()
}

func ExampleTask_WithHandler() {
	Do(func(t time.Time) error {
		fmt.Println("executing scheduled task")
		return fmt.Errorf("expected error")
	}).
		WithHandler(func(errs <-chan error) {
			for err := range errs {
				fmt.Println("failed: ", err.Error())
			}
		}).
		After(5 * time.Second).
		Schedule().
		Wait()
}

func ExampleTask_Every() {
	Do(func(t time.Time) error {
		fmt.Println("executing scheduled task")
		return nil
	}).
		Every(5 * time.Second).
		Schedule().
		CancelWhen(time.After(30 * time.Second)).
		Wait()
}
