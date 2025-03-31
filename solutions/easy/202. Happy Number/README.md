**Problem**: Happy Number
**Link**: [LeetCode Problem](https://leetcode.com/problems/happy-number/)

**Complexity**:

- **Time**: O(log n), where n is the given number. Each iteration processes the digits of `n`, which takes logarithmic time in base 10.
- **Space**: O(log n), since we use a hash map to track seen numbers, which at most stores `log n` values before detecting a cycle or reaching 1.

**Explanation**:
This function determines whether a number is a happy number. A happy number is defined as a number that, when replaced by the sum of the squares of its digits repeatedly, eventually reaches 1. If the process enters a cycle, the number is not happy. The function utilizes a hash map to track previously seen numbers to detect cycles efficiently.

**Performance**:

- **Runtime (LeetCode)**: 0ms (100%)
