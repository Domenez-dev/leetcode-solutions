Problem: Add Two Numbers  
Link: https://leetcode.com/problems/add-two-numbers/

Complexity:

- Time: O(max(m, n)), where m and n are the lengths of the two linked lists.
- Space: O(max(m, n)), for the resulting linked list.

Runtime on LeetCode:

- 0ms (Beats 100%)

Approach:

1. **Traverse Both Lists**: Iterate through both linked lists, summing corresponding digits along with any carry from the previous step.
2. **Handle Carry**: Calculate the carry and the current digit (`sum % 10`). [Learn more about carry in addition](<https://en.wikipedia.org/wiki/Carry_(arithmetic)>).
3. **Build the Result List**: Use a dummy node to simplify list construction, appending nodes for each sum.
4. **Edge Cases**:
   - Lists of different lengths.
   - Remaining carry after traversal.
   - Input lists with leading zeros (e.g., `0 -> 0 -> 1`).
