# Problem: Maximum Value of an Ordered Triplet I

**Link:** [LeetCode Problem](https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/)

## Complexity:

- **Time Complexity:** O(n³), where n is the number of elements in the array. The function uses three nested loops to iterate through all possible triplets.
- **Space Complexity:** O(1), as the function only uses a few integer variables for computation.

## Explanation:

This function calculates the maximum triplet value from an array using the formula **(nums[i] - nums[j]) \* nums[k]**, where **i < j < k** and **nums[j] < nums[i]**.

The function iterates over all possible triplets using three nested loops. It checks if **nums[j]** is smaller than **nums[i]** before computing the triplet value. The maximum value found is returned as the result.

## Performance:

- **Runtime (LeetCode):** 0ms(100%)
