# Write your MySQL query statement below
SELECT c.name "Customers" FROM Customers c WHERE c.id NOT IN (SELECT o.customerId FROM Orders o) ;
