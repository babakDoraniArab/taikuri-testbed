---
spec: SPEC-002
section: orders
version: 1
deps: [SPEC-001]
---

# Orders

## Requirements

- R1: Every order SHALL reference an existing customer.
- R2: Order persistence SHALL go through the orders repository, not ad-hoc SQL.
