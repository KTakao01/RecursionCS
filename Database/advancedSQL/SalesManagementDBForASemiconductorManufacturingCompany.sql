-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1059
SELECT c.id 会社id,c.name 会社名,c.contact_email 会社連絡,SUM(p.quantity*p.item_price) 合計取引額 
FROM purchases p INNER JOIN companies c ON p.company_id = c.id 
GROUP BY c.id ORDER BY c.id;