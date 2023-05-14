package golang

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	oddHead, evenHead, evenFirstNode := head, head.Next, head.Next
	for oddHead.Next != nil && evenHead.Next != nil {
		oddHead.Next = oddHead.Next.Next
		evenHead.Next = evenHead.Next.Next
		oddHead = oddHead.Next
		evenHead = evenHead.Next
	}
	oddHead.Next = evenFirstNode
	return head
}
