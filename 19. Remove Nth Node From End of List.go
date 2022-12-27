package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var lastSlow *ListNode
	slow := head
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	if slow == fast && n == 1 {
		return nil
	}
	for {
		if fast == nil {
			if n == 1 {
				lastSlow.Next = nil
			} else {
				if lastSlow == nil {
					head = head.Next
				} else {
					lastSlow.Next = lastSlow.Next.Next
				}
			}
			return head
		}
		lastSlow = slow
		slow = slow.Next
		fast = fast.Next
	}
	return head
}
