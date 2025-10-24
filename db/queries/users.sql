INSERT INTO users (name, dob)
VALUES (?, ?);
SELECT id, name, dob FROM users WHERE id= ?;
UPDATE users
SET name = ?, dob =?
WHERE id =?;
DELETE FROM users where id=?;
SELECT id,name,dob FROM users;