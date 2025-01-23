Problem: String to Integer (atoi)  
Link: https://leetcode.com/problems/string-to-integer-atoi/

Complexity:

- Time: O(n), where n is the length of the string.
- Space: O(1), constant space usage.
- Runtime (LeetCode): 0ms (100.00%)

Approach:

1. **Skip Whitespace**: Traverse the string to ignore leading spaces.
2. **Handle Sign**: Determine the sign based on '+' or '-' and set the `positive` flag accordingly.
3. **Convert to Integer**: Use a map to convert characters to digits while checking for overflows.
4. **Clamp Values**: Ensure the result stays within the 32-bit integer range.

Edge Cases:

- Empty string or no digits.
- Input exceeding 32-bit integer range.
- Non-numeric characters or invalid input format.
