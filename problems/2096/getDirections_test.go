package exercises2096

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindsAndBuildsPaths(t *testing.T) {
	Convey("TestFindsAndBuildsPaths", t, func() {
		root := &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}
		result := getDirections(root, 2, 3)
		So(result, ShouldEqual, "UR")
	})
	// [5,1,2,3,null,6,4]
	Convey("TestFindsAndBuildsPaths", t, func() {
		root := &TreeNode{Val: 5,
			Left: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
				Right: nil,
			},
			Right: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
				Right: &TreeNode{Val: 4, Left: nil, Right: nil},
			},
		}
		result := getDirections(root, 3, 6)
		So(result, ShouldEqual, "UURL")
	})
}
