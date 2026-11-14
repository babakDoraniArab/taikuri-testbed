# migrations

Sequential SQL migrations for the testbed fixture, applied in filename order.

| File | Creates |
|------|---------|
| `0001_customers.sql` | `customers` (`id`, `name`, `email`) |
| `0002_orders.sql` | `orders` — `customer_id` is a FK `REFERENCES customers(id)` |

## Plantable violation site — referenced-column drop

VIOLATION-SITE: column-drop — a future migration `0003_*.sql` doing
`ALTER TABLE customers DROP COLUMN email;` — or, worse, dropping `customers.id`,
which `orders.customer_id` references — is the fixture for the M7 ER-based check
`migration-drops-referenced-column`. Do NOT add that migration yet; this note
marks where the referenced-column-drop test will plant it.
