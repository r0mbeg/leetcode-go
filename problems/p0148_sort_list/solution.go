package p0148_sort_list

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	isSorted := false

	for !isSorted {

		cur := head

		for cur.Next != nil {
			if cur.Val > cur.Next.Val {

				//fmt.Printf("%d > %d, swapping \n", cur.Val, cur.Next.Val)

				isSorted = false
				buffer := cur.Next.Val
				cur.Next.Val = cur.Val
				cur.Val = buffer
				cur = head
				continue
			}
			cur = cur.Next
		}

		isSorted = true

	}

	return head

}

func PrintList(head *ListNode) string {
	if head == nil {
		return "<nil>"
	}

	res := ""
	cur := head

	for cur != nil {
		res += strconv.Itoa(cur.Val)
		if cur.Next != nil {
			res += " -> "
		}
		cur = cur.Next
	}

	return res
}
