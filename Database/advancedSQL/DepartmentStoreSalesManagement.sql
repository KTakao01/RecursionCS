-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1045
SELECT COUNT(price) ,MAX(price),MIN(price),AVG(price),SUM(price)
FROM items WHERE sold_at IS NOT NULL AND sold_at LIKE "2021-06-%";
