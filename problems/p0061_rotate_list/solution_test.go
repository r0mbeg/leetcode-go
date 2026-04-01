package p0061_rotate_list

import (
	"fmt"
	"testing"
)

func TestRotateList(t *testing.T) {

	type tc struct {
		list ListNode
		k    int
		want ListNode
	}

	//node3 := &ListNode{2, nil}
	//node2 := &ListNode{1, node3}
	//node1 := &ListNode{1, node2}

	cases := []tc{

		{
			list: ListNode{1, nil},
			k:    1,
			want: ListNode{1, nil},
		},

		{
			list: ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			k:    4,
			want: ListNode{2, &ListNode{0, &ListNode{1, nil}}},
		},
	}

	for i, c := range cases {

		got := rotateRight(&c.list, c.k)
		if !AreEqualList(*got, c.want) {
			t.Fatalf("case %d : rotateRight(%s, %d) = %s; want %s", i, PrintList(&c.list), c.k, PrintList(got), PrintList(&c.want))
		}
		fmt.Printf("case %d correct : rotateRight(%s, %d) = %s \n", i, PrintList(&c.list), c.k, PrintList(got))
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
