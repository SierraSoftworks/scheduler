package strat

import (
	"fmt"
	"testing"

	"time"

	"sync"

	. "github.com/smartystreets/goconvey/convey"
)

func TestManual(t *testing.T) {
	Convey("Manual", t, func() {
		Convey("Strategy", func() {
			s := Manual()

			Convey("String()", func() {
				So(s.String(), ShouldEqual, "manual")
			})

			Convey("Schedule()", func() {
				sc := s.Schedule()
				So(sc, ShouldNotBeNil)
				So(sc, ShouldHaveSameTypeAs, &manualSchedule{})

				So(s.targets, ShouldHaveLength, 1)

				sc.Cancel()
				So(s.targets, ShouldHaveLength, 0)
			})
		})

		Convey("Schedule", func() {
			s := Manual()
			So(s, ShouldNotBeNil)
			sc := s.Schedule()
			So(sc, ShouldNotBeNil)
			So(s.targets, ShouldHaveLength, 1)

			Convey("Events()", FailureContinues, func(c C) {
				select {
				case <-sc.Events():
					So(fmt.Errorf("triggered"), ShouldBeNil)
				default:
				}

				t := time.Time{}

				var wgReady sync.WaitGroup
				wgReady.Add(1)
				var wgDone sync.WaitGroup
				wgDone.Add(1)
				go func() {
					wgReady.Done()
					select {
					case t = <-sc.Events():
					case <-time.After(5 * time.Millisecond):
						c.So(fmt.Errorf("not triggered"), ShouldBeNil)
					}

					wgDone.Done()
				}()

				wgReady.Wait()

				s.Trigger()

				wgDone.Wait()

				So(t, ShouldHappenWithin, 5e6*time.Nanosecond, time.Now())
			})

			Convey("Cancel()", func() {
				sc.Cancel()
				So(s.targets, ShouldHaveLength, 0)
			})

		})
	})
}
