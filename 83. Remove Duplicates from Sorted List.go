package golang

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curVal := head.Val
	lastNode := head
	nextNode := head.Next
	for nextNode != nil {
		if nextNode.Val != curVal {
			curVal = nextNode.Val
			lastNode.Next = nextNode
			lastNode = lastNode.Next
			nextNode = nextNode.Next
			continue
		}
		if nextNode.Val == curVal && nextNode.Next == nil {
			lastNode.Next = nil
		}
		nextNode = nextNode.Next
	}
	return head
}
