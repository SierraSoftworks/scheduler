package strat

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterval(t *testing.T) {
	Convey("Interval", t, func() {
		s := Every(1 * time.Second)

		Convey("Strategy", func() {
			Convey("String()", func() {
				So(s.String(), ShouldEqual, "every 1s")
			})

			Convey("Interval()", func() {
				So(s.Interval(), ShouldEqual, time.Second)
			})

			Convey("Schedule()", func() {
				So(s.Schedule(), ShouldNotBeNil)
				So(s.Schedule(), ShouldHaveSameTypeAs, &intervalSchedule{})
			})
		})

		Convey("Schedule", func() {
			sc := newIntervalSchedule(time.Millisecond)
			So(sc, ShouldNotBeNil)

			Convey("Events()", func() {
				e := sc.Events()
				So(e, ShouldNotBeNil)

				t1 := <-e
				t2 := <-e

				So(t2, ShouldHappenAfter, t1)
				So(t2, ShouldHappenOnOrBetween, t1.Add(-3e6*time.Nanosecond), t1.Add(3e6*time.Nanosecond))
			})

			Convey("Cancel()", func() {
				sc.Cancel()

				_, ok := <-sc.Events()
				So(ok, ShouldBeFalse)
			})
		})
	})
}
