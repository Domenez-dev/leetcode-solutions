Problem: Find Median of Two Sorted Arrays
Link: [median-of-two-sorted-arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

Complexity:

- Time: O(log(min(m, n)))
- Space: O(1)

Runtime on LeetCode:

- 0ms (Beats 100%)

Approach:

1. **Binary Search**: Perform binary search on the smaller array to find the correct partition of the combined arrays. [Learn more about binary search](https://en.wikipedia.org/wiki/Binary_search_algorithm).
2. **Partitions**: Divide the arrays into left and right halves ensuring the maximum value in the left half is less than or equal to the minimum value in the right half.
3. **Median Calculation**:
   - If the total length is even, return the average of the max left and min right values.
   - If the total length is odd, return the max of the left partition.
4. **Edge Cases**: Handle cases where partitions are at the start or end of the arrays using `math.MinInt32` and `math.MaxInt32`.

If no valid partition is found (which theoretically shouldn't happen for sorted arrays), the program will panic.
