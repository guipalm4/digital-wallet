USE wallet;

CREATE TABLE IF NOT EXISTS customers (id varchar(255), name varchar(255), email varchar(255), created_at date, updated_at date);
CREATE TABLE IF NOT EXISTS accounts (id varchar(255), customer_id varchar(255), balance decimal(12, 2), created_at date);
CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount decimal(12, 2), created_at date);
