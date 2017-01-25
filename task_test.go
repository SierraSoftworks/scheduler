package scheduler

import (
	"fmt"
	"testing"
	"time"

	"github.com/SierraSoftworks/scheduler/strat"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Task(t *testing.T) {
	Convey("Task", t, func() {

		Convey("New()", func() {
			x := New()
			So(x, ShouldNotBeNil)
			So(x.Action(), ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Handler(), ShouldNotBeNil)
		})

		Convey("Do", func() {
			f := func(t time.Time) error {
				return nil
			}
			x := Do(f)
			So(x, ShouldNotBeNil)
			So(x.Action(), ShouldEqual, f)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Never())

			Convey("Nill Action", func() {
				x := Do(nil)
				So(x, ShouldNotBeNil)
				So(x.Action(), ShouldNotBeNil)
			})
		})

		Convey("Instance", func() {
			x := New()
			So(x, ShouldNotBeNil)
			So(x.Action(), ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)

			Convey("Clone()", func() {
				xa := x.Action()
				xs := x.Strategy()

				c := x.Clone()
				So(c, ShouldNotBeNil)
				So(c, ShouldNotEqual, x)
				So(c.Action(), ShouldEqual, x.Action())
				So(c.Strategy(), ShouldEqual, x.Strategy())

				So(x.Action(), ShouldEqual, xa)
				So(x.Strategy(), ShouldEqual, xs)
			})

			Convey(".String()", func() {
				So(x.String(), ShouldEqual, "Task(never)")
				So(x.Immediately().String(), ShouldEqual, "Task(immediate)")
			})

			Convey("Do()", func() {
				xa := x.Action()
				xs := x.Strategy()

				f := func(t time.Time) error {
					return nil
				}
				y := x.Do(f)

				So(x.Action(), ShouldEqual, xa)
				So(x.Strategy(), ShouldEqual, xs)

				So(y, ShouldNotBeNil)
				So(y, ShouldNotEqual, x)
				So(y.Action(), ShouldEqual, f)

				Convey("Nill Action", func() {
					y := x.Do(nil)
					So(y, ShouldNotBeNil)
					So(y.Action(), ShouldNotBeNil)
				})
			})

			Convey("WithStrategy()", func() {
				xa := x.Action()
				xs := x.Strategy()

				y := x.WithStrategy(strat.Immediate())
				So(y, ShouldNotBeNil)
				So(y.Strategy(), ShouldHaveSameTypeAs, strat.Immediate())
				So(y.Strategy().String(), ShouldEqual, "immediate")

				So(x.Action(), ShouldEqual, xa)
				So(x.Strategy(), ShouldEqual, xs)

				Convey("Nil Strategy", func() {
					y := x.WithStrategy(nil)
					So(y, ShouldNotBeNil)
					So(y.Strategy(), ShouldNotBeNil)
					So(y.Strategy(), ShouldHaveSameTypeAs, strat.Never())
				})
			})

			Convey("WithHandler()", func() {
				xa := x.Action()
				xs := x.Strategy()
				xh := x.Handler()

				f := func(c <-chan error) {
					for err := range c {
						fmt.Println("Handled Error: ", err.Error())
					}
				}
				y := x.WithHandler(f)

				So(y, ShouldNotBeNil)
				So(y.Handler(), ShouldEqual, f)

				So(x.Action(), ShouldEqual, xa)
				So(x.Strategy(), ShouldEqual, xs)
				So(x.Handler(), ShouldEqual, xh)

				Convey("Nill Handler", func() {
					y := x.WithHandler(nil)
					So(y, ShouldNotBeNil)
					So(y.Handler(), ShouldNotBeNil)
				})
			})

			Convey("Schedule()", func(c C) {
				ran := false
				ranArg := time.Time{}
				complete := make(chan struct{})

				tBefore := time.Now()
				s := Do(func(t time.Time) error {
					ran = true
					ranArg = t
					return fmt.Errorf("complete")
				}).
					WithStrategy(strat.Immediate()).
					WithHandler(func(errs <-chan error) {
						for err := range errs {
							c.So(err, ShouldNotBeNil)
							c.So(err.Error(), ShouldEqual, "complete")
						}

						complete <- struct{}{}
					}).
					Schedule()

				So(s, ShouldNotBeNil)

				select {
				case <-complete:
					So(ran, ShouldBeTrue)
					So(ranArg, ShouldHappenOnOrBetween, tBefore, time.Now())
					So(s.LastError(), ShouldNotBeNil)

				case <-time.After(10 * time.Millisecond):
					So(fmt.Errorf("timed out"), ShouldBeNil)
				}
			})
		})
	})
}
