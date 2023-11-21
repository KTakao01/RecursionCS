-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1049
SELECT user_id AS ユーザーID,COUNT(ticket_id) AS 予約回数 FROM reservations GROUP BY ユーザーID HAVING 予約回数 >= 10;
