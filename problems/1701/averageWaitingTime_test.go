package exercises1701

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAverageWaitingTime(t *testing.T) {
	Convey("TestAverageWaitingTime", t, func() {
		Convey("1", func() {
			result := AverageWaitingTime([][]int{{1, 2}, {2, 5}, {4, 3}})
			So(result, ShouldEqual, 5.0)
		})

		Convey("2", func() {
			result := AverageWaitingTime([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}})
			So(result, ShouldEqual, 3.25)
		})
	})
}
