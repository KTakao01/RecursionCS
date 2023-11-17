-- https://recursionist.io/dashboard/course/6/lesson/1033
-- クエリを書いてください。 --
SELECT * FROM expenses WHERE spent >= 200 AND is_asset = 0 AND (type IN ("entertainment", "outside_dining", "travel_vacation", "shopping")OR type IS NULL);

--SELECT * FROM expenses WHERE spent >= 200 AND is_asset = 0 AND type IN ("entertainment", "outside_dining", "travel_vacation", "shopping", IS NULL);
-- IN句内でIS NULLを使うことはできません。IN句は、指定された値のリストの中から値を比較するために使用されますが、NULLとの比較は特別な扱いを要します。