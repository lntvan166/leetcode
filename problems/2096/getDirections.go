// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/?envType=daily-question&envId=2024-07-16
package exercises2096

import "strings"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	s := &strings.Builder{}
	d := &strings.Builder{}
	find(root, startValue, s)
	find(root, destValue, d)
	i := 0
	maxI := min(d.Len(), s.Len())
	for i < maxI && s.String()[s.Len()-i-1] == d.String()[d.Len()-i-1] {
		i++
	}
	return strings.Repeat("U", s.Len()-i) + reverse(d.String())[i:]
}

func find(n *TreeNode, val int, sb *strings.Builder) bool {
	if n.Val == val {
		return true
	}
	if n.Left != nil && find(n.Left, val, sb) {
		sb.WriteString("L")
	} else if n.Right != nil && find(n.Right, val, sb) {
		sb.WriteString("R")
	}
	return sb.Len() > 0
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
