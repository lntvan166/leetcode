package exercises1636

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFrequencySortAscendingOrder(t *testing.T) {
	Convey("TestFrequencySortAscendingOrder", t, func() {
		Convey("Case with mixed frequencies", func() {
			nums := []int{4, 4, 1, 1, 1, 2, 2, 3}
			result := frequencySort(nums)
			expected := []int{3, 4, 4, 2, 2, 1, 1, 1}
			So(result, ShouldResemble, expected)
		})
	})
}
