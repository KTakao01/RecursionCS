-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1040
UPDATE comments SET content = 'You are so talented. I do hope these research papers are studied more so they can someday be implemented in real life!' WHERE id = 23;
UPDATE comments SET type = 'reply' WHERE type = 'replied';
SELECT * FROM comments WHERE type = 'reply' ORDER BY id LIMIT 30;
