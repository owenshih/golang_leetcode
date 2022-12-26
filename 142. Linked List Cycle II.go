package golang

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	smap := make(map[*ListNode]int)
	haveCycle := -1
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for {
		smap[slow] += 1
		if fast == nil || fast.Next == nil {
			break
		}
		if slow == fast {
			haveCycle = 0
			slow = slow.Next
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	if haveCycle == -1 {
		return nil
	} else {
		for {
			_, ok := smap[slow]
			if ok {
				return slow
			}
			smap[slow] += 1
			slow = slow.Next
		}
		return nil
	}
}
