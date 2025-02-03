INSERT INTO customers (id, name, email, created_at, updated_at)
VALUES
    ('c7f47b44-bd18-4a3b-913f-48a8f99cdb4f', 'John Doe', 'john@example.com', '2025-01-01', '2025-02-03'),
    ('6124f3ee-8671-4ba9-a3a5-8124425bb29c', 'Jane Roe', 'jane@example.com', '2025-01-01', '2025-02-03');

INSERT INTO accounts (id, customer_id, balance, created_at)
VALUES
    ('31c7ec79-5a0a-47bc-bc8d-b7d3f2072605', 'c7f47b44-bd18-4a3b-913f-48a8f99cdb4f', 2000.00, '2025-02-03'),
    ('fe69c70b-4096-4160-977b-86ab3ac1c9fa', '6124f3ee-8671-4ba9-a3a5-8124425bb29c', 2000.00, '2025-02-03');