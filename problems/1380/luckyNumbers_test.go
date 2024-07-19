package exercises1380

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Returns correct lucky numbers for a given matrix
func TestReturnsCorrectLuckyNumbers(t *testing.T) {
	Convey("TestReturnsCorrectLuckyNumbers", t, func() {
		Convey("Matrix with one lucky number", func() {
			matrix := [][]int{
				{3, 7, 8},
				{9, 11, 13},
				{15, 16, 17},
			}
			result := luckyNumbers(matrix)
			So(result, ShouldResemble, []int{15})
		})

		Convey("Matrix with multiple lucky numbers", func() {
			matrix := [][]int{
				{1, 10, 4, 2},
				{9, 3, 8, 7},
				{15, 16, 17, 12},
			}
			result := luckyNumbers(matrix)
			So(result, ShouldResemble, []int{12})
		})
	})
}
