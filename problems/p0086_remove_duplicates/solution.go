package p0086_remove_duplicates

// ListNode is a definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}
