// https://leetcode.com/problems/swap-nodes-in-pairs/submissions/965208052/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head.Next
	prev := (*ListNode)(nil)
	node := head
	for node != nil && node.Next != nil {
		next := node.Next.Next
		node.Next.Next = node
		if prev != nil {
			prev.Next = node.Next
		}
		node.Next = next
		prev = node
		node = next
	}
	return ans
}
