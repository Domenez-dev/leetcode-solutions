**Problem**: Path Sum  
**Link**: [LeetCode Problem](https://leetcode.com/problems/path-sum/)

**Complexity**:

- **Time**: O(n), where n is the number of nodes in the tree. The function traverses each node once to check if a path with the target sum exists.
- **Space**: O(h), where h is the height of the tree. The space is used for recursion stack.

**Explanation**:  
This function checks whether the binary tree has a root-to-leaf path such that the sum of the node values equals the given `targetSum`. It uses a depth-first search approach by recursively traversing the left and right subtrees. The base case checks if the current node is a leaf (both `Left` and `Right` are `nil`) and if the remaining `targetSum` matches the node's value.

**Performance**:

- **Runtime (LeetCode)**: 15ms (3.86%)
