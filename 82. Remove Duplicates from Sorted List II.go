package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	var headNode, mainNode, lastNode *ListNode
	nextNode := head
	lastVal := -101
	invalidVal := -101
	for nextNode != nil {
		if lastVal == nextNode.Val {
			invalidVal = nextNode.Val
		}
		if nextNode.Val != lastVal && lastVal != invalidVal {
			if mainNode == nil {
				headNode = lastNode
			}
			if mainNode != nil {
				mainNode.Next = lastNode
			}
			mainNode = lastNode
			mainNode.Next = nil
		}
		lastNode = nextNode
		lastVal = nextNode.Val
		nextNode = nextNode.Next
	}
	if lastVal != invalidVal {
		if mainNode == nil {
			headNode = lastNode
		}
		if mainNode != nil {
			mainNode.Next = lastNode
		}
		mainNode = lastNode
		mainNode.Next = nil
	}
	return headNode
}
