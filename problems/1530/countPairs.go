// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/description/?envType=daily-question&envId=2024-07-18
package exercises1530

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// todo review
// leetcode solution
func countPairs(root *TreeNode, distance int) int {
	count := 0
	const MaxDistance = 10

	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return make([]int, MaxDistance+1)
		}

		if node.Left == nil && node.Right == nil {
			res := make([]int, MaxDistance+1)
			res[1] = 1
			return res
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		for i := 1; i <= distance; i++ {
			for j := 1; j <= distance-i; j++ {
				count += left[i] * right[j]
			}
		}

		res := make([]int, MaxDistance+1)
		for i := 1; i < MaxDistance; i++ {
			res[i+1] = left[i] + right[i]
		}

		return res
	}

	dfs(root)
	return count
}
