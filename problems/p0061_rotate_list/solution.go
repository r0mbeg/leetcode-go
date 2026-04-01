package p0061_rotate_list

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	//fmt.Println("length of ", PrintList(head), " is", length)

	// resolve cycle
	k = k % length
	if k == 0 {
		return head
	}

	tail.Next = head // cycle

	stepsToNewTail := length - k - 1
	newTail := head
	for i := 0; i < stepsToNewTail; i++ {
		newTail = newTail.Next
	}

	// делаем разрыв
	newHead := newTail.Next
	newTail.Next = nil

	return newHead

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
