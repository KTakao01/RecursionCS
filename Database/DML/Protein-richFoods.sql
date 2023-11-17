-- https://recursionist.io/dashboard/course/6/lesson/1034
-- クエリを書いてください。 --
SELECT 
    food_name, 
    carbohydrates, 
    protein, fat, 
    (carbohydrates * 4 + protein * 4 + fat * 9) AS total_calories 
FROM foods 
WHERE protein >= 40
AND carbohydrates < 20
AND fat < 10;
