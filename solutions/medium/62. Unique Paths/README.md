### Problem: Unique Paths

**Link**: [https://leetcode.com/problems/unique-paths/](https://leetcode.com/problems/unique-paths/)

**Complexity**:

- **Time**: O(m \* n), `m`: number of rows and `n`: number of columns.
- **Space**: O(m \* n), the space required for memoization map `cache` and the recursion stack.
- **Runtime (LeetCode)**: 0 ms (100%)

**Explanation**:

1. The problem is solved recursively with memoization (Dynamic Programming.
2. The function `moving` calculates the number of unique paths from the top-left to the bottom-right of the grid.
   - If the robot reaches the bottom-right corner (`m == 1 && n == 1`), return `1` (one valid path).
   - If the robot moves out of bounds (`m == 0 || n == 0`), return `0` (no valid path).
   - For each valid state `(m, n)`, compute the sum of paths by moving right `(m-1, n)` and down `(m, n-1)`.
3. Results are cached in the `cache` map using keys `(m, n)` and `(n, m)` to avoid recomputing symmetric states.
4. The final result is the total number of unique paths.

**Key Insights**:

- Memoization is critical to avoid redundant recursive calls, reducing the time complexity from exponential to polynomial.

**Resources**:

- [Dynamic Programming Guide, geeksforgeeks](https://www.geeksforgeeks.org/dynamic-programming/)
- [Dynamic Programming - Learn to Solve Algorithmic Problems & Coding Challenges](https://www.youtube.com/watch?v=oBt53YbR9Kk&t=4588s)
