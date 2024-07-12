package exercises1598

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCorrectOperationCount(t *testing.T) {
	Convey("TestCorrectOperationCount", t, func() {
		Convey("1", func() {
			logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
			result := minOperations(logs)
			So(result, ShouldEqual, 3)
		})
		Convey("2", func() {
			logs := []string{"./", "../", "./"}
			result := minOperations(logs)
			So(result, ShouldEqual, 0)
		})
		Convey("3", func() {
			logs := []string{"d1/", "d2/", "../", "d21/", "./"}
			result := minOperationsLeetCode(logs)
			So(result, ShouldEqual, 2)
		})
	})
}
