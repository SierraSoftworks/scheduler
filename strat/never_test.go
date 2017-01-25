package strat

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNever(t *testing.T) {
	Convey("Never", t, func() {
		Convey("Strategy", func() {
			s := Never()
			Convey("String()", func() {
				So(s.String(), ShouldEqual, "never")
			})

			Convey("Schedule()", func() {
				sc := s.Schedule()
				So(sc, ShouldNotBeNil)
				So(sc, ShouldHaveSameTypeAs, &neverSchedule{})
			})
		})

		Convey("Schedule", func() {
			sc := newNeverSchedule()
			So(sc, ShouldNotBeNil)

			Convey("Events()", func() {
				e := sc.Events()
				So(e, ShouldNotBeNil)

				_, ok := <-e
				So(ok, ShouldBeFalse)
			})

			Convey("Cancel()", func() {
				sc.Cancel()

				_, ok := <-sc.Events()
				So(ok, ShouldBeFalse)
			})
		})
	})
}
