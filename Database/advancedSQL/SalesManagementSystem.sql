-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1050
SELECT
 member_id,COUNT(*) 
FROM
 negotiations 
WHERE 
 status = 'SUCCESS' AND done_at LIKE '2021-10-%' 
GROUP BY
 member_id 
ORDER BY 
 COUNT(*) DESC,
 member_id;
