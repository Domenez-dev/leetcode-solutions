## Problem: Valid Arrangement of Pairs

**Link:** [https://leetcode.com/problems/valid-arrangement-of-pairs/](https://leetcode.com/problems/valid-arrangement-of-pairs/)

### Complexity:

- **Time:** O(N), where N is the number of edges (pairs). This is because we traverse each edge once during Eulerian path construction.
- **Space:** O(N), due to the adjacency list representation of the graph and the recursion stack in the worst case.

### Runtime on LeetCode:

- **730ms Beats 60.71%%**

### Approach:

1. **Graph Construction:**

   - Use an adjacency list to represent the directed graph.
   - Maintain `in-degree` and `out-degree` counts for each node.

2. **Finding the Start Node:**

   - Identify the node where `out-degree > in-degree` as the starting node (Eulerian path property).
   - If no such node exists, start from any node in the graph.

3. **Hierholzer’s Algorithm for Eulerian Path:**

   - Use a stack-based DFS to traverse edges while ensuring Eulerian path conditions.
   - Backtrack when no more outgoing edges exist for a node.
   - Reverse the obtained path to get the correct order.

4. **Constructing the Result:**
   - Convert the Eulerian path into a list of ordered pairs representing the valid arrangement.

### Edge Cases:

- The given pairs already form a cycle.
- Disconnected components (invalid input for this problem's constraints).
- A single pair (trivial case).
- Large inputs to test efficiency.

This approach ensures an optimal and correct reconstruction of a valid arrangement using **Hierholzer’s Algorithm** for Eulerian paths in directed graphs.
