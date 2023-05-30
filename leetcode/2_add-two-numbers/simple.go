/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{}
	node := start
	for {
		l1Next := (*ListNode)(nil)
		l2Next := (*ListNode)(nil)
		if l1 != nil {
			node.Val += l1.Val
			l1Next = l1.Next
		}
		if l2 != nil {
			node.Val += l2.Val
			l2Next = l2.Next
		}
		if l1Next == nil && l2Next == nil && node.Val < 10 {
			break
		}
		node.Next = &ListNode{Val: node.Val / 10}
		node.Val %= 10
		node = node.Next
		l1 = l1Next
		l2 = l2Next
	}
	return start
}
