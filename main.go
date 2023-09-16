package main

import (
	"poivre/listnode"
)

func addTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	res := &listnode.ListNode{}
	ret := res
	more := 0

	for {
		current := l1.Val + l2.Val + more
		more = 0
		if current >= 10 {
			res.Val = current - 10
			more = 1
		} else {
			res.Val = current
		}

		if l1.Next == nil && l2.Next == nil {
			if more == 1 {
				res.Next = &listnode.ListNode{
					Val: 1,
				}
			}
			return ret
		}
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}

		res.Next = &listnode.ListNode{}
		res = res.Next
	}
}
