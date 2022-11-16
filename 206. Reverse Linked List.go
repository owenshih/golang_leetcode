package golang

func reverseList(head *ListNode) *ListNode {
	var prevNode, tmpNode *ListNode
	currNode := head
	for currNode != nil {
		tmpNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = tmpNode
	}
	return prevNode
}
