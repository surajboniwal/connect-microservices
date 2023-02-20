SELECT * FROM users
WHERE email = $1 LIMIT 1;

INSERT INTO users (
  fullname, username, email, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;