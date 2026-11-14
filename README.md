# taikuri-testbed

Shared test fixture for [taikuri](https://github.com/dotech/taikuri). Every taikuri
demo and conformance-gate test in the milestone plan runs against this repo.

It is deliberately small and deliberately "plantable": a few well-known sites are
marked with `VIOLATION-SITE:` comments where later milestones inject real
violations to prove the gate catches them. **No live violations are planted yet** —
this repo builds clean (`go build ./...` passes).

## Layout

| Path | What |
|------|------|
| `cmd/svc/main.go` | Tiny `net/http` service; imports `internal/orders`. |
| `internal/customers/` | Customer domain model + service. |
| `internal/orders/` | Order model + repo; **imports `internal/customers`** (the internal import edge). |
| `frontend/` | Dependency-free TypeScript stub. |
| `migrations/` | Sequential SQL migrations — `customers`, then `orders` with a FK. |

## Plantable violation sites

| Site | File | Marker / planted violation |
|------|------|----------------------------|
| Raw SQL from input | `internal/orders/repo.go` | `VIOLATION-SITE: raw SQL from input` → `query := "SELECT * FROM orders WHERE id=" + id` |
| Hardcoded secret | `internal/customers/service.go` | `VIOLATION-SITE: hardcoded secret` → `apiKey := "sk_live_..."` |
| Cross-package import | `cmd/svc/main.go` | `CROSS-PACKAGE-IMPORT SITE` — chain `cmd/svc → internal/orders → internal/customers`; plant a forbidden edge here |
| Referenced-column drop | `migrations/README.md` | `VIOLATION-SITE: column-drop` — a `0003_*.sql` dropping a referenced column |

Each milestone uncomments / adds exactly the violation it needs, runs the gate,
then reverts — keeping `main` green between demos.
