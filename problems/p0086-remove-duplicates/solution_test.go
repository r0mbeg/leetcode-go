package p0086_remove_duplicates

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {

	type tc struct {
		list ListNode
		want ListNode
	}

	node3 := &ListNode{2, nil}
	node2 := &ListNode{1, node3}
	node1 := &ListNode{1, node2}

	cases := []tc{
		{
			list: ListNode{},
			want: ListNode{},
		},

		{
			list: ListNode{1, nil},
			want: ListNode{1, nil},
		},

		{
			list: *node1,
			want: ListNode{1, &ListNode{2, nil}},
		},
	}

	for _, c := range cases {
		got := deleteDuplicates(&c.list)
		if calcDepth(got) != calcDepth(&c.want) {
			t.Fatalf("deleteDuplicates(%v) = %v; want %v", c.list, got, c.want)
		}
		gotSlice := listToSlice(got)
		wantSlice := listToSlice(&c.want)

		if !reflect.DeepEqual(gotSlice, wantSlice) {
			t.Fatalf("deleteDuplicates(%v) = %v; want %v", c.list, gotSlice, wantSlice)
		}

	}
}

func calcDepth(list *ListNode) int {
	l := list
	depth := 0
	for l != nil {
		depth++
		l = l.Next
	}
	return depth
}

func listToSlice(l *ListNode) []int {
	res := make([]int, 0)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}
