**Problem**: Find the Champion  
**Link**: [LeetCode Problem](https://leetcode.com/problems/find-the-champion/)

**Complexity**:

- **Time**: O(n), where n is the number of participants (or grid size). The function iterates through each row to determine the champion.
- **Space**: O(1), as the function uses only a constant amount of space.

**Explanation**:  
This function finds the champion in a competition represented by a 2D grid (`grid`), where `grid[i][j] = 1` means participant `i` defeated participant `j`. The function uses a two-pointer approach to identify a candidate champion. It iterates through each participant and checks whether the candidate defeats others or is defeated. If the candidate is defeated, the current participant becomes the new candidate. The final candidate is returned as the champion.

**Performance**:

- **Runtime (LeetCode)**: 0ms (100%)
