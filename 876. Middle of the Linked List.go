package golang

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	slowNode, fastNode := head, head
	var slowIndex, fastIndex float64 = 0, 0
	for fastNode != nil {
		fastNode = fastNode.Next
		fastIndex++
		var middleCount = math.Floor(fastIndex / 2)
		if middleCount > slowIndex {
			slowNode = slowNode.Next
			slowIndex++
		}
	}
	return slowNode
}
