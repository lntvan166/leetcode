package exercises1334

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Correctly identifies the city with the smallest number of reachable cities within the distance threshold
func TestFindTheCityWithSmallestReachableCities(t *testing.T) {
	Convey("TestFindTheCityWithSmallestReachableCities", t, func() {
		Convey("Example 1", func() {
			n := 4
			edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
			distanceThreshold := 4
			result := findTheCity(n, edges, distanceThreshold)
			So(result, ShouldEqual, 3)
		})

		Convey("Example 2", func() {
			n := 5
			edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
			distanceThreshold := 2
			result := findTheCity(n, edges, distanceThreshold)
			So(result, ShouldEqual, 0)
		})
	})
}
