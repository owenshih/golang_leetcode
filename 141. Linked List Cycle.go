package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	var nextArray []*ListNode
	for head != nil {
		var inArray = inArray(head, nextArray)
		if inArray {
			return true
		}
		nextArray = append(nextArray, head)
		head = head.Next
	}
	return false
}

func inArray(node *ListNode, array []*ListNode) bool {
	for _, v := range array {
		if v == node {
			return true
		}
	}
	return false
}
