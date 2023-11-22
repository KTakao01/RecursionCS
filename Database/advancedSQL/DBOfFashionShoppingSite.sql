-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1060
SELECT c.id category_id,c.name category_name,SUM(new_pro.sales_price) total_sales_amount
FROM
 (
    SELECT pu.id id,pu.product_id product_id,pro.product_id product_id,pu.sales_price sales_price,pro.category_id category_id
    FROM purchases pu
    INNER JOIN product_categories pro
    ON pu.product_id = pro.product_id 
) new_pro
INNER JOIN
 categories c
ON new_pro.category_id = c.id
GROUP BY c.id
ORDER BY total_sales_amount DESC;
