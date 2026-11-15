---
spec: SPEC-001
section: customers
version: 1
deps: []
---

# Customers

## Requirements

- R1: Customer records SHALL be retrieved only through the customers service layer.
- R2: Customer lookups SHALL use parameterised queries, never string-built SQL.
- R3: Service credentials SHALL be loaded from the environment, never committed to source.
