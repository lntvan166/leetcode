package exercises1530

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountPairsBalancedTree(t *testing.T) {
	Convey("TestCountPairsBalancedTree", t, func() {
		Convey("Balanced Tree with distance 3", func() {
			root := &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			}
			result := countPairs(root, 3)
			So(result, ShouldEqual, 2)
		})
	})
	Convey("TestCountPairsNullRoot", t, func() {
		Convey("Null root node", func() {
			var root *TreeNode = nil
			result := countPairs(root, 3)
			So(result, ShouldEqual, 0)
		})
	})
}
