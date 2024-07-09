package exercises2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	i1, i2 := l1, l2
	temp := 0
	head := &ListNode{}
	current := head
	for i1 != nil || i2 != nil {
		sum := temp
		if i1 != nil {
			sum += i1.Val
			i1 = i1.Next
		}
		if i2 != nil {
			sum += i2.Val
			i2 = i2.Next
		}
		temp = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if temp > 0 {
		current.Next = &ListNode{Val: temp}
	}

	return head.Next
}
