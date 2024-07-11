package exercises2

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNumbers(t *testing.T) {
	// 2 -> 4 -> 3
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	// 5 -> 6 -> 4
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := AddTwoNumbers(l1, l2)
	Convey("TestAddTwoNumbers", t, func() {
		So(result.Val, ShouldEqual, 7)
		So(result.Next.Val, ShouldEqual, 0)
		So(result.Next.Next.Val, ShouldEqual, 8)
	})
}
