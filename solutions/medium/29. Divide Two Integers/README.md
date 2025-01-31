### Problem: Divide Two Integers

**Link**: [https://leetcode.com/problems/divide-two-integers/](https://leetcode.com/problems/divide-two-integers/)

**Complexity**:

- **Time**: O(n), where n is the absolute value of `dividend`. The loop subtracts `divisor` repeatedly, making it inefficient for large numbers.
- **Space**: O(1), as only a few variables are used.
- **Runtime (LeetCode)**: 1184 ms (15%)

**Notes**:

- To get 0ms runtime n leetcode you must use the division / which triggers the contraint in the description:

```
 divide two integers without using multiplication, division, and mod operato
```

- This implementation handles overflow by capping results at `2^31 - 1` when needed.
