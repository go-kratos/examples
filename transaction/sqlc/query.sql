-- name: CreateUser :execresult
INSERT INTO users (name) VALUES (?);

-- name: CreateUserDetail :execresult
INSERT INTO user_details (id,email) VALUES (?,?);
