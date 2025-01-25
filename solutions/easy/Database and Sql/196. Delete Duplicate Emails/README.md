**Problem**: Delete Duplicate Records Based on Email  
**Link**: [LeetCode Problem](https://leetcode.com/problems/delete-duplicate-emails/)

**Query**:

```sql
DELETE FROM Person WHERE id NOT IN (
    SELECT id FROM (
        SELECT MIN(id) "id" FROM Person GROUP BY email
    ) AS UniqueEmail
)
```

**Explanation**:  
This query deletes duplicate records from the `Person` table, keeping only the record with the smallest `id` for each unique `email`. The query uses a subquery to identify the minimum `id` for each email and then deletes all other records that have the same email but a higher `id`.

**Complexity**:

- **Time**: O(n), where n is the number of rows in the `Person` table. The query performs a full table scan to identify duplicates.
- **Space**: O(n), where n is the number of rows in the `Person` table. Space is used to store the result of the subquery that identifies the minimum `id` for each email.
- **Runtime**: 860ms (80%)
  This query effectively removes duplicate entries based on the email field while preserving the record with the smallest `id`. Let me know if you'd like to modify or add anything!
