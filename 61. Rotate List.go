package golang

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	nodeArray := make([]*ListNode, 0)
	for head != nil {
		nodeArray = append(nodeArray, head)
		head = head.Next
	}
	nodeLen := len(nodeArray)
	if nodeLen == 0 {
		return head
	}
	rotateCount := k % nodeLen
	if rotateCount == 0 {
		return nodeArray[0]
	}
	head = nodeArray[nodeLen-rotateCount]
	nodeArray[nodeLen-rotateCount-1].Next = nil
	nodeArray[nodeLen-1].Next = nodeArray[0]
	return head
}
