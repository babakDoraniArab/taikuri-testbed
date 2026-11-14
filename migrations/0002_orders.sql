-- 0002_orders.sql — orders table with a foreign key into customers.
CREATE TABLE orders (
    id          TEXT PRIMARY KEY,
    customer_id TEXT NOT NULL REFERENCES customers(id),
    total_cents INTEGER NOT NULL
);
