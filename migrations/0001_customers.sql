-- 0001_customers.sql — base customers table.
CREATE TABLE customers (
    id    TEXT PRIMARY KEY,
    name  TEXT NOT NULL,
    email TEXT NOT NULL
);
