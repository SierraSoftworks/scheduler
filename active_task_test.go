package scheduler

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActiveTask(t *testing.T) {
	Convey("ActiveTask", t, func() {
		t := New()
		So(t, ShouldNotBeNil)

		Convey("New()", func() {
			a := newActiveTask(t)
			So(a, ShouldNotBeNil)
			So(a.task, ShouldEqual, t)
			So(a.schedule, ShouldHaveSameTypeAs, t.Strategy().Schedule())
			So(a.errors, ShouldNotBeNil)
			So(a.lastError, ShouldBeNil)
		})

		Convey("Instance", func() {

			Convey("String()", func() {
				a := newActiveTask(t)
				So(a, ShouldNotBeNil)

				So(a.String(), ShouldEqual, "Active Task(never)")
			})

			Convey("Wait()", func() {
				hasExecuted := false

				t := t.
					Do(func(t time.Time) error {
						hasExecuted = true
						return nil
					}).
					After(10 * time.Millisecond)

				a := newActiveTask(t)
				So(a, ShouldNotBeNil)

				t0 := time.Now()
				a.run()
				So(hasExecuted, ShouldBeFalse)

				a.Wait()
				t1 := time.Now()

				So(t1, ShouldHappenWithin, 2*time.Millisecond, t0.Add(10*time.Millisecond))
				So(hasExecuted, ShouldBeTrue)
			})

			Convey("Done()", func() {
				hasExecuted := false

				t := t.
					Do(func(t time.Time) error {
						hasExecuted = true
						return nil
					}).
					After(10 * time.Millisecond)

				a := newActiveTask(t)
				So(a, ShouldNotBeNil)

				t0 := time.Now()
				a.run()
				So(hasExecuted, ShouldBeFalse)

				d := a.Done()
				t1 := <-d
				So(t1, ShouldHappenWithin, time.Millisecond, t0.Add(10*time.Millisecond))
				So(hasExecuted, ShouldBeTrue)

				_, ok := <-d
				So(ok, ShouldBeFalse)
			})

			Convey("Cancel()", func() {
				hasExecuted := false

				t := t.
					Do(func(t time.Time) error {
						hasExecuted = true
						return nil
					}).
					After(10 * time.Millisecond)

				a := newActiveTask(t)
				So(a, ShouldNotBeNil)

				a.run()
				So(hasExecuted, ShouldBeFalse)

				a.Cancel()
				a.Wait()

				So(hasExecuted, ShouldBeFalse)
			})

			Convey("CancelWhen()", func() {
				hasExecuted := false

				t := t.
					Do(func(t time.Time) error {
						hasExecuted = true
						return nil
					}).
					After(10 * time.Millisecond)

				a := newActiveTask(t)
				So(a, ShouldNotBeNil)

				Convey("Nill Channel", func() {
					a.run()
					So(hasExecuted, ShouldBeFalse)

					So(a.CancelWhen(nil), ShouldEqual, a)

					a.Wait()
					So(hasExecuted, ShouldBeTrue)
				})

				Convey("Closed Channel", func() {
					a.run()
					So(hasExecuted, ShouldBeFalse)

					c := make(chan time.Time)
					close(c)
					So(a.CancelWhen(c), ShouldEqual, a)

					a.Wait()
					So(hasExecuted, ShouldBeTrue)
				})

				Convey("Valid Channel", func() {
					a.run()
					So(hasExecuted, ShouldBeFalse)

					So(a.CancelWhen(time.After(1*time.Millisecond)), ShouldEqual, a)

					a.Wait()
					So(hasExecuted, ShouldBeFalse)
				})
			})

			Convey("LastError()", func() {
				a := newActiveTask(t.Do(func(t time.Time) error {
					return fmt.Errorf("expected failure")
				}).Immediately())

				a.run()
				a.Wait()
				So(a.LastError(), ShouldNotBeNil)
				So(a.LastError().Error(), ShouldEqual, "expected failure")
			})

			Convey("Run", func() {

			})
		})
	})
}
