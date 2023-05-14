package golang

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	sumVal, carry := 0, 0
	for l1 != nil || l2 != nil {
		sumVal = 0
		if l2 != nil {
			sumVal += l2.Val
			l2 = l2.Next
		}
		if l1 != nil {
			sumVal += l1.Val + carry
			if sumVal >= 10 {
				carry = 1
			} else {
				carry = 0
			}
			l1.Val = sumVal % 10
			if l1.Next == nil && l2 == nil && carry == 1 {
				l1.Next = &ListNode{Val: 1}
				break
			}
			if l1.Next == nil && l2 != nil {
				l1.Next, l2 = l2, l1.Next
			}
			l1 = l1.Next
		}
	}
	return head
}
