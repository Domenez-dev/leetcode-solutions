/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	var prev *ListNode

	for head != nil && head.Next != nil {
		second := head.Next
		nextPair := second.Next
		second.Next = head
		head.Next = nextPair

		if prev != nil {
			prev.Next = second
		}
		prev = head
		head = nextPair
	}
	return newHead
}
