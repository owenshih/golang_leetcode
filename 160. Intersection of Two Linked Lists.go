package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var linkANodeArray, linkBNodeArray []*ListNode
	var linkA, linkB = headA, headB
	for linkA != nil || linkB != nil {
		if linkA != nil {
			linkANodeArray = append(linkANodeArray, linkA)
			var isANode = inArray(linkA, linkBNodeArray)
			if isANode {
				return linkA
			}
			linkA = linkA.Next
		}
		if linkB != nil {
			linkBNodeArray = append(linkBNodeArray, linkB)
			var isBNode = inArray(linkB, linkANodeArray)
			if isBNode {
				return linkB
			}
			linkB = linkB.Next
		}
	}
	return nil

}

func inArray(node *ListNode, nodeArray []*ListNode) bool {
	for _, v := range nodeArray {
		if v == node {
			return true
		}
	}
	return false
}
