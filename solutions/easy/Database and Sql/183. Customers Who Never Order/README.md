**Problem**: Customers Who Never Ordered  
**Link**: [LeetCode Problem](https://leetcode.com/problems/customers-who-never-ordered/)

**Complexity**:

- **Time**: O(n + m), where n is the number of customers and m is the number of orders. The query requires scanning the `Customers` table once and the `Orders` table once to identify customers who have not placed any orders.
- **Space**: O(n), where n is the number of customers. This is the space required to store the results of the query.

**Performance**:

- **Runtime (LeetCode)**: 504ms (95%)
