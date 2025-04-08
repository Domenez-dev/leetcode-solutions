Problem: 3396. Minimum Number of Operations to Make Elements in Array Distinct

Link: [Similar LeetCode Problem](https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct)

Time: O(n)
Space: O(n)

Explanation:
Counts element frequencies, then recursively removes the first 3 elements until remaining elements are distinct or fewer than 3 remain.

Performance:

    Runtime: ~7ms

    Memory: ~8.6MB

Note: Each operation removes the first 3 elements. Handles edge cases (n < 3) optimally.
