# Write your MySQL query statement below

SELECT d.name AS "Department", e.name AS "Employee", e.salary AS "Salary"
FROM Employee e
JOIN Department d ON d.id = e.departmentId
WHERE (e.departmentId, e.salary) IN
    (SELECT e.departmentId, MAX(e.salary)
     FROM Employee e
     GROUP BY e.departmentId);

