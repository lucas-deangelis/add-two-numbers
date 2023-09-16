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

func addTwoNumbersCodeium(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	res := &listnode.ListNode{}
	ret := res
	more := 0

	for l1 != nil || l2 != nil {
		current := more
		if l1 != nil {
			current += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current += l2.Val
			l2 = l2.Next
		}

		more = current / 10
		res.Val = current % 10

		if l1 != nil || l2 != nil || more != 0 {
			res.Next = &listnode.ListNode{}
			res = res.Next
		}
	}

	if more != 0 {
		res.Next = &listnode.ListNode{
			Val: 1,
		}
	}

	return ret
}

func addTwoNumbersGPT4(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	dummy := &listnode.ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		cur.Next = &listnode.ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummy.Next
}

func addTwoNumbersLeetcode(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	result := &listnode.ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &listnode.ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &listnode.ListNode{}
		}
		tmp = tmp.Next
	}
	return result
}
