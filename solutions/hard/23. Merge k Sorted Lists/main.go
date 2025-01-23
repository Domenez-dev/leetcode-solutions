/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	head := &ListNode{}
	current := head

	for minHeap.Len() > 0 {
		smallest := heap.Pop(minHeap).(*ListNode)
		current.Next = smallest
		current = current.Next

		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}

	return head.Next
}
