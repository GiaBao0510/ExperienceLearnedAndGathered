-- SQL schema for User and Order tables
CREATE TABLE users (
    uuid VARCHAR(255) PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(10),
    password VARCHAR(255) NOT NULL
);

CREATE TABLE orders (
    order_id VARCHAR(255) PRIMARY KEY,
    uuid VARCHAR(255) NOT NULL,
    order_date DATE NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (uuid) REFERENCES users(uuid)
);