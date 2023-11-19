-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1039
DELETE FROM posts WHERE id IN (1,5,6) OR category = 'travel_vacation';
SELECT * FROM posts ORDER BY category,title;
