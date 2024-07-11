package exercises4

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMedianSortedArrays(t *testing.T) {
	result := findMedianSortedArrays([]int{1, 3}, []int{2})
	Convey("TestFindMedianSortedArrays", t, func() {
		So(result, ShouldEqual, 2.0)
	})
}
