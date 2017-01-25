package strat

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestImmediate(t *testing.T) {
	Convey("Immediate", t, func() {
		Convey("Strategy", func() {
			Convey("String()", func() {
				s := Immediate()
				So(s, ShouldNotBeNil)
				So(s.String(), ShouldEqual, "immediate")
			})

			Convey("Schedule()", func() {
				s := Immediate()
				So(s, ShouldNotBeNil)

				sc := s.Schedule()
				So(sc, ShouldNotBeNil)
				So(sc, ShouldHaveSameTypeAs, &immediateSchedule{})
			})
		})

		Convey("Schedule", func() {
			sc := newImmediateSchedule()
			So(sc, ShouldNotBeNil)

			Convey("Events()", func() {
				e := sc.Events()
				So(e, ShouldNotBeNil)

				_, ok := <-e
				So(ok, ShouldBeTrue)
			})

			Convey("Cancel()", func() {
				sc.Cancel()

				_, ok := <-sc.Events()
				So(ok, ShouldBeFalse)
			})
		})
	})
}
