package scheduler

import (
	"testing"
	"time"

	"github.com/SierraSoftworks/scheduler/strat"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStrategies(t *testing.T) {
	Convey("Strategies", t, func() {
		x := New()
		So(x, ShouldNotBeNil)

		Convey("At", func() {
			x := x.At(time.Now().Add(30 * time.Second))
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.At(time.Now()))
		})

		Convey("After", func() {
			x := x.After(time.Second)
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Delay(time.Second))
		})

		Convey("Every", func() {
			x := x.Every(time.Second)
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Every(time.Second))
		})

		Convey("Never", func() {
			x := x.Never()
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Never())
		})

		Convey("Manual", func() {
			s := strat.Manual()
			x := x.Manual(s)
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Manual())
		})

		Convey("Immediately", func() {
			x := x.Immediately()
			So(x, ShouldNotBeNil)
			So(x.Strategy(), ShouldNotBeNil)
			So(x.Strategy(), ShouldHaveSameTypeAs, strat.Immediate())
		})
	})
}
