Problem: Partition Equal Subset Sum  
Link: https://leetcode.com/problems/partition-equal-subset-sum/

Complexity:

    Time: O(n * sum/2), where n is the number of elements and sum is the total sum of the array.
    Space: O(sum/2) for the DP array.
    Runtime (LeetCode): 10ms (86.59%)

Explanation:

The goal is to determine if the array can be split into two subsets with equal sums. First, we calculate the total sum of the array. If the sum is odd, it's impossible to split it equally, so we return false immediately.

If the sum is even, we aim to find a subset whose sum equals `sum / 2`. We use a dynamic programming (DP) approach where `dp[i]` represents whether a subset with sum `i` can be formed.

We initialize `dp[0] = true`, since a sum of 0 can always be formed by selecting no elements.

Then for each number in the array, we iterate backwards through the DP array (from `target` down to `num`) and update `dp[j]` to `true` if `dp[j - num]` was already `true`. This prevents using the same number multiple times in one iteration.

Finally, we return `dp[target]`, which tells us if a subset with the target sum exists.
