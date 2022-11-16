package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var headNode, lastNode *ListNode
	currentNode := head
	for currentNode != nil {
		if currentNode.Val != val {
			if headNode == nil {
				headNode = currentNode
				lastNode = headNode
			} else {
				lastNode.Next = currentNode
				lastNode = currentNode
			}
		}
		if currentNode.Next == nil && currentNode.Val == val && lastNode != nil {
			lastNode.Next = nil
		}
		currentNode = currentNode.Next
	}
	return headNode
}
