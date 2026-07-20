-- Tệp tin query.sql này chủ yếu chứa các câu lệnh thực thi SQL

-- name: CreateUsers :exec
INSERT INTO users (uuid, user_name, email, phone_number, password) VALUES ($1, $2, $3, $4, $5);
-- name: GetUsersByID :one
SELECT * FROM users WHERE uuid = $1;
-- name: GetAllUserss :many
SELECT * FROM users;

-- name: CreateOrder :exec
INSERT INTO orders (order_id, uuid, order_date, total_amount, status) VALUES ($1, $2, $3, $4, $5);
-- name: GetOrderByID :one
SELECT * FROM orders WHERE order_id = $1;
-- name: GetAllOrders :many
SELECT * FROM orders;
