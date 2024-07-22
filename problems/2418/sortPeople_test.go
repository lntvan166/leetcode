package exercises2418

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortPeopleByHeightDescending(t *testing.T) {
	Convey("TestSortPeopleByHeightDescending", t, func() {
		Convey("Sort people by height", func() {
			names := []string{"Alice", "Bob", "Charlie"}
			heights := []int{150, 180, 165}
			result := sortPeople(names, heights)
			expected := []string{"Bob", "Charlie", "Alice"}
			So(result, ShouldResemble, expected)
		})
	})
	Convey("TestSortPeopleByHeightDescending", t, func() {
		Convey("Sort people by height", func() {
			names := []string{"Alice", "Bob", "Charlie", "David"}
			heights := []int{150, 180, 165, 160}
			result := sortPeople(names, heights)
			expected := []string{"Bob", "Charlie", "David", "Alice"}
			So(result, ShouldResemble, expected)
		})
	})
}
