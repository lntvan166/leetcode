package exercises1190

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseParentheses(t *testing.T) {
	Convey("TestReverseParentheses", t, func() {
		Convey("1", func() {
			result := ReverseParentheses("(abcd)")
			So(result, ShouldEqual, "dcba")
		})

		Convey("2", func() {
			result := ReverseParentheses("(uoy evol I)")
			So(result, ShouldEqual, "I love you")
		})

		Convey("3", func() {
			result := ReverseParentheses("(ed(et(oc))el)")
			So(result, ShouldEqual, "leetcode")
		})

		Convey("4", func() {
			result := ReverseParentheses("a(bcdefghijkl(mno)p)q")
			So(result, ShouldEqual, "apmnolkjihgfedcbq")
		})
	})
}
