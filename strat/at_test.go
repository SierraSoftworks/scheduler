package strat

import (
	"fmt"
	"testing"

	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAt(t *testing.T) {
	Convey("At", t, func() {

		Convey("Strategy", func() {
			Convey("String()", func() {
				epoc, err := time.Parse("2006-02-01T15:04:05Z", "1970-01-01T00:00:00Z")
				So(err, ShouldBeNil)

				s := At(epoc)
				So(s, ShouldNotBeNil)

				So(s.String(), ShouldEqual, "at 1970-01-01 00:00:00 +0000 UTC")
				So(s.String(), ShouldEqual, fmt.Sprintf("at %s", s.t.String()))
			})

			Convey("Time()", func() {
				t := time.Now()
				s := At(t)
				So(s, ShouldNotBeNil)
				So(s.Time(), ShouldResemble, t)
			})

			Convey("Schedule()", func() {
				Convey("Before Now", func() {
					s := At(time.Now().Add(-5 * time.Second))
					So(s, ShouldNotBeNil)
					sc := s.Schedule()
					So(sc, ShouldNotBeNil)
					So(sc, ShouldHaveSameTypeAs, &neverSchedule{})
				})

				Convey("After Now", func() {
					s := At(time.Now().Add(5 * time.Second))
					So(s, ShouldNotBeNil)
					sc := s.Schedule()
					So(sc, ShouldNotBeNil)
					So(sc, ShouldHaveSameTypeAs, &delaySchedule{})
				})
			})
		})
	})
}
