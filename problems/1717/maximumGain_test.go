package exercises1717

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaximumGain(t *testing.T) {
	Convey("TestMaximumGain", t, func() {
		Convey("1", func() {
			result := maximumGainLeetCode("cdbcbbaaabab", 4, 5)
			So(result, ShouldEqual, 19)
		})

		Convey("2", func() {
			result := maximumGain("aabbaaxybbaabb", 5, 4)
			So(result, ShouldEqual, 20)
		})

		Convey("3", func() {
			result := maximumGainLeetCode("ababaab", 1, 1)
			So(result, ShouldEqual, 3)
		})
	})
}
