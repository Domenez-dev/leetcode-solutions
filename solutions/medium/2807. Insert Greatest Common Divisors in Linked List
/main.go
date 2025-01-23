/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	dummy := head
	for dummy.Next != nil {
		var newNode ListNode
		newNode.Val = pgcd(dummy.Val, dummy.Next.Val)
		newNode.Next = dummy.Next
		dummy.Next = &newNode
		dummy = newNode.Next
	}
	return head
}

func pgcd(n, m int) int {
	if m == 0 {
		return n
	} else {
		return pgcd(m, n%m)
	}
}
