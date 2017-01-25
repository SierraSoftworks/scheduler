package strat

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDelay(t *testing.T) {
	Convey("Delay", t, func() {

		Convey("Strategy", func() {
			Convey("String()", func() {
				s := Delay(time.Second)
				So(s.String(), ShouldEqual, "after 1s")
			})

			Convey("Delay()", func() {
				s := Delay(time.Second)
				So(s.Delay(), ShouldEqual, time.Second)
			})

			Convey("Schedule()", func() {
				Convey("Negative Delay", func() {
					s := Delay(-1 * time.Second)
					sc := s.Schedule()
					So(sc, ShouldNotBeNil)
					So(sc, ShouldHaveSameTypeAs, &neverSchedule{})
				})

				Convey("Positive Delay", func() {
					s := Delay(time.Second)
					sc := s.Schedule()
					So(sc, ShouldNotBeNil)
					So(sc, ShouldHaveSameTypeAs, &delaySchedule{})
				})
			})

		})

		Convey("Schedule", func() {
			t0 := time.Now()
			sc := newDelaySchedule(time.Millisecond)

			Convey("Events()", func() {
				t := <-sc.Events()
				So(t, ShouldHappenOnOrBetween, t0, t0.Add(3e6*time.Nanosecond))

				_, ok := <-sc.Events()
				So(ok, ShouldBeFalse)
			})

			Convey("Cancel()", func() {
				sc.Cancel()
				_, ok := <-sc.Events()
				So(ok, ShouldBeFalse)

				Convey("Subsequent Calls", func() {
					sc.Cancel()
					_, ok := <-sc.Events()
					So(ok, ShouldBeFalse)
				})
			})

		})

	})
}
