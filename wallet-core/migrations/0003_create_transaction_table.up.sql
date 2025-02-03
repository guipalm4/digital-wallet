CREATE TABLE IF NOT EXISTS transactions (
    id varchar(255) PRIMARY KEY ,
    account_id_from varchar(255) NOT NULL ,
    account_id_to varchar(255) NOT NULL ,
    amount decimal(12, 2),
    created_at date
);
