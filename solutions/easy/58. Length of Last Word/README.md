### Problem: Length of Last Word

**Link**: [https://leetcode.com/problems/length-of-last-word/](https://leetcode.com/problems/length-of-last-word/)

**Complexity**:

- **Time**: O(n), where `n` is the length of the input string `s`. The loop iterates through the string once, starting from the end.
- **Space**: O(1), as no additional space is used other than a few variables.
- **Runtime (LeetCode)**: 0 ms (100%)

**Explanation**:

1. Start iterating from the end of the string `s`.
2. If a non-space character is found, increment the `lenght` counter and set the `in_word` flag to `true`.
3. If a space character is encountered after `in_word` becomes `true`, return the accumulated `lenght`.
4. If the end of the string is reached without encountering another space, return the total `lenght`.
