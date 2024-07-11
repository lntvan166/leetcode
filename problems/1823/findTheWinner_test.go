package exercises1823

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindTheWinner(t *testing.T) {
	Convey("TestFindTheWinner", t, func() {
		Convey("1", func() {
			result := FindTheWinner(5, 2)
			So(result, ShouldEqual, 3)
		})

		Convey("2", func() {
			result := FindTheWinner(6, 5)
			So(result, ShouldEqual, 1)
		})

		Convey("3", func() {
			result := FindTheWinner(7, 7)
			So(result, ShouldEqual, 5)
		})
	})
}
