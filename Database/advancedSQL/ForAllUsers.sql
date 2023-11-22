-- クエリを書いてください。 --
-- https://recursionist.io/dashboard/course/6/lesson/1054
ALTER TABLE customers ADD COLUMN type VARCHAR(255) DEFAULT 'customers';
ALTER TABLE leads ADD COLUMN type VARCHAR(255) DEFAULT 'leads';
ALTER TABLE employees ADD COLUMN type VARCHAR(255) DEFAULT 'employees';
ALTER TABLE leads ADD COLUMN username VARCHAR(255);

SELECT username,email,customer_first_name first_name,customer_last_name last_name,created_at,type FROM customers
UNION
SELECT username,email,first_name,last_name,created_at,type FROM leads
UNION
SELECT username,email,employee_first_name first_name,employee_last_name last_name,created_at,type FROM employees 
ORDER BY created_at DESC;
