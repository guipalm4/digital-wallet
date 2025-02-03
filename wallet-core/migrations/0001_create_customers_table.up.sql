CREATE TABLE IF NOT EXISTS customers (
    id varchar(255) PRIMARY KEY NOT NULL,
    name varchar(255) NOT NULL ,
    email varchar(255) NOT NULL UNIQUE ,
    created_at date,
    updated_at date
);