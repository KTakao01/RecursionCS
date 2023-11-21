-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1048
SELECT tag_id,COUNT(content) AS 投稿数 FROM posts GROUP BY tag_id ORDER BY 投稿数 DESC LIMIT 30;
