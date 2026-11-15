---
spec: SPEC-001
section: customers
version: 3
deps: []
---

# Customers

## Requirements

- R1: Customer records SHALL be retrieved only through the customers service layer; direct repository access from handlers is forbidden.
- R2: Customer lookups SHALL use parameterised queries exclusively; dynamic SQL strings are forbidden.
- R3: Service credentials SHALL be loaded from the environment, never committed to source.
