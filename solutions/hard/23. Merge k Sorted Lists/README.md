Problem: Merge k Sorted Lists  
Link: https://leetcode.com/problems/merge-k-sorted-lists/

Complexity:

- Time: O(N log k), N is the total number of nodes and k is the number of linked lists.
- Space: O(k).

Runtime on LeetCode:

- 0ms (Beats 100%)

Approach:

1. **Min-Heap**: Use a priority queue (min-heap) to efficiently fetch the smallest node among the heads of the `k` lists. [Learn more about heaps](<https://en.wikipedia.org/wiki/Heap_(data_structure)>).
2. **Heap Operations**:
   - Initialize the heap with the head of each non-empty list.
   - Pop the smallest node from the heap and add it to the result.
   - If the popped node has a next node, push it to the heap.
3. **Linked List Construction**: Build the merged linked list by iteratively connecting the smallest nodes.

Edge Cases:

- All input lists are empty.
- Only one list is non-empty.
- Lists with duplicate values.
