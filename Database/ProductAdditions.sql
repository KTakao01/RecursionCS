-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1038
INSERT INTO products (product_name,price,brand,model,description,category) VALUES 
('32GB RAM DDR5',150,'Corsair','Vengeance','Fast speed 2 stick ram DDR5 technology','electronics'),
('Harry Potter and the Sorcerer Stone I',10,'Bloomsbury',NULL,'A book about Harry Potter and his magical journey.','books'),
('Ceiling Fan',350,'Minka','Aire Xtreme H2O','A coal wet ceiling fan with remote control. It has 8 different spin speeds.','home');

SELECT * FROM products ORDER BY product_name;

