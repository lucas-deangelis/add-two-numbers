package main

import (
	"math/rand"
	"poivre/listnode"
	"reflect"
	"testing"
	"testing/quick"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *listnode.ListNode
		l2 *listnode.ListNode
	}
	tests := []struct {
		name string
		args args
		want *listnode.ListNode
	}{
		{
			name: "342 + 465 = 807",
			args: args{
				l1: listnode.FromInt(342),
				l2: listnode.FromInt(465),
			},
			want: listnode.FromInt(807),
		},
		{
			name: "0 + 0 = 0",
			args: args{
				l1: listnode.FromInt(0),
				l2: listnode.FromInt(0),
			},
			want: listnode.FromInt(0),
		},
		{
			name: "9999999 + 9999 = 10009998",
			args: args{
				l1: listnode.FromInt(9999999),
				l2: listnode.FromInt(9999),
			},
			want: listnode.FromInt(10009998),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddListNodes(t *testing.T) {
	f := func(a, b *listnode.ListNode) bool {
		result := addTwoNumbers(a, b)
		expected := addListNodesBruteForce(a, b)

		if !reflect.DeepEqual(result, expected) {
			t.Logf("Failed case:\n  a: %v,\n  b: %v, res: %v,\nexp: %v", a, b, result, expected)
			return false
		}

		return reflect.DeepEqual(result, expected)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func addListNodesBruteForce(a, b *listnode.ListNode) *listnode.ListNode {
	res := a.ToInt() + b.ToInt()
	return listnode.FromInt(res)
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	l1 := listnode.FromInt(rand.Int())
	l2 := listnode.FromInt(rand.Int())

	b.Run("base", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addTwoNumbers(l1, l2)
		}
	})

	b.Run("codeium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addTwoNumbersCodeium(l1, l2)
		}
	})

	b.Run("GPT-4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addTwoNumbersGPT4(l1, l2)
		}
	})

	b.Run("Leetcode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addTwoNumbersLeetcode(l1, l2)
		}
	})
}
