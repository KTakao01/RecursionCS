-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1051
SELECT
 product_id,
 warehouse_id,
 CASE 
  WHEN month IN (12,1,2,3) THEN year || '-' || printf("%02d",month)
  ELSE NULL
 END AS stock_year_month,
 COUNT(*) AS all_count
FROM warehouse
WHERE month IN (12,1,2,3)
GROUP BY product_id,warehouse_id,stock_year_month 
ORDER BY year DESC,month DESC,warehouse_id,product_id;
