package p0148_sort_list

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {

	type testcase struct {
		list ListNode
		want ListNode
	}

	cases := []testcase{
		{
			list: ListNode{4, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
			want: ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		},

		{
			list: ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}},
			want: ListNode{-1, &ListNode{0, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},

		{
			list: ListNode{},
			want: ListNode{},
		},
	}

	for i, c := range cases {
		got := sortList(&c.list)

		if !AreEqualList(*got, c.want) {
			t.Fatalf("case %d incorrect : sortList(%v) = (%v); want (%v)", i, PrintList(&c.list), PrintList(got), PrintList(&c.want))
		}

		fmt.Printf("case %d correct : sortList(%v) = (%v) \n", i, PrintList(&c.list), PrintList(got))
	}

}

func AreEqualList(a, b ListNode) bool {
	for a.Next != nil && b.Next != nil {
		if a.Next.Val != b.Next.Val {
			return false
		}
		a.Next = a.Next.Next
		b.Next = b.Next.Next
	}
	return a == b
}
