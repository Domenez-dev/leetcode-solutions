**Problem**: Matching Phone Numbers  
**Link**: [LeetCode Problem](https://leetcode.com/problems/valid-phone-number/)

**Command**:

```bash
grep -Ex "[0-9]{3}-[0-9]{3}-[0-9]{4}|\([0-9]{3}\) [0-9]{3}-[0-9]{4}" file.txt
```

**Explanation**:  
This `grep` command is used to match valid phone numbers in the following formats:

- `XXX-XXX-XXXX`, where `X` is a digit.
- `(XXX) XXX-XXXX`, where `X` is a digit and the area code is enclosed in parentheses.

**Complexity**:

- **Time**: O(n), where n is the number of lines in the file. The command scans each line of the file to check if it matches the regular expression.
- **Space**: O(1), as the command uses a constant amount of space for matching patterns.
- **Runtime**: 65ms (55%).
