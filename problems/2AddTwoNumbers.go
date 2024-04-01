package leetcode

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// create 2 nodes - 1 is to preserve, 2 will be result.
	// dummy's head is empty, the result will be next node.
	dummy := new(ListNode)
	l3 := dummy
	var sum int
	var carry int
	// Walking over list 'till both won't be exhausted.
	for l1 != nil || l2 != nil {
		// Check if first node is nil, take it's value if not.
		var x int
		if l1 != nil {
			x = l1.Val
		}
		// Check if second node is nil, take it's value if not.
		var y int
		if l2 != nil {
			y = l2.Val
		}
		// Add 2 nodes vals and carry.
		sum = x + y + carry
		// If carry more than 9, divide by 10 and keep for next iteration.
		carry = sum / 10
		// In the new node there will be remainder of sum.
		// E.g. if it's first iteration x = 6 and y = 7, then 6 + 7 + 0 = 13.
		// It's over than 9 -> carry is 1, next node will hold 4.
		l3.Next = &ListNode{Val: sum % 10}
		// Update l3 to be next node of new node to repeat operation.
		l3 = l3.Next
		// Update 1 and 2 nodes.
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// if last node was bigger than 9, it should produce carry which will be next node.
	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
