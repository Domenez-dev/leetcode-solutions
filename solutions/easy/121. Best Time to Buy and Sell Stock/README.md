**Problem**: Best Time to Buy and Sell Stock  
**Link**: [LeetCode Problem](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

**Complexity**:

- **Time**: O(n²), where n is the number of elements in the `prices` array. The function uses two nested loops to find the highest price after each potential low price, resulting in a quadratic time complexity.
- **Space**: O(1), as the function only uses a few variables (`profit`, `high`, and `low`) to store intermediate results.

**Explanation**:  
This function aims to calculate the maximum profit from a series of stock prices, where you can buy at one price and sell at a later price. It loops through the array and keeps track of the lowest price (`low`) and the highest price after that low (`high`). For each potential buying price, it computes the profit and keeps track of the maximum profit seen so far. The function uses a helper function `max` to return the larger of two values.

**Performance**:

- **Runtime (LeetCode)**: 198ms (5%)
