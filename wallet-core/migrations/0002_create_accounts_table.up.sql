CREATE TABLE IF NOT EXISTS accounts (
    id varchar(255) PRIMARY KEY ,
    customer_id varchar(255),
    balance decimal(12, 2),
    created_at date
);