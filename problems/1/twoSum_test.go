package exercises1

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTwoSum(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	Convey("TestTwoSum", t, func() {
		So(result, ShouldResemble, []int{0, 1})
	})
}
