package listnode

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromInt(n int) *ListNode {
	s := numberToSlice(n)
	slices.Reverse(s)
	return fromSlice(s)
}

func (l *ListNode) ToInt() int {
	total := 0
	pow := 0

	for {
		total += l.Val * int(math.Pow10(pow))
		pow++
		if l.Next == nil {
			return total
		}
		l = l.Next
	}
}

// Thanks GPT-4
func numberToSlice(n int) []int {
	str := strconv.Itoa(n)
	slice := make([]int, len(str))
	for i, r := range str {
		digit, _ := strconv.Atoi(string(r))
		slice[i] = digit
	}
	return slice
}

func fromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	ret := res
	for i := 1; i < len(nums); i++ {
		res.Next = &ListNode{
			Val: nums[i],
		}
		res = res.Next
	}
	return ret
}

func toSlice(l *ListNode) []int {
	res := []int{}
	for {
		if l == nil {
			return res
		}
		res = append(res, l.Val)
		l = l.Next
	}
}

func (l *ListNode) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(FromInt(size))
}

// Thanks GPT-4
func (n *ListNode) String() string {
	var s []string
	for current := n; current != nil; current = current.Next {
		s = append(s, fmt.Sprintf("%d", current.Val))
	}
	return strings.Join(s, " -> ")
}
