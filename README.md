# Order-to-Cash & Settlement Fabric

Global multi-entity order-to-cash platform: concurrent reservation, double-entry ledger, multi-currency settlement, multilateral netting, and verifiable settlement receipts.

## Status

Scaffold. Domain model, reservation core, ledger, and settlement engine are not implemented yet.

| Stage | Scope | State |
|-------|-------|-------|
| 0 | Service skeleton, config, health, CI-ready build | done |
| 1 | Single-warehouse reservation + order state machine | next |
| 2 | Double-entry ledger + inventory costing | planned |
| 3 | Multi-warehouse / cross-entity reservation | planned |
| 4 | Order-to-invoice saga | planned |
| 5 | Netting + cryptographic receipts | planned |
| 6 | Credit and policy engine | planned |
| 7 | Regional sync + hardening | planned |

## Documents

- [Product overview](10-order-to-cash-settlement.md)
- [PRD](10-d01-prd.md)
- [Architecture](docs/architecture.md)

## Quick start

```sh
go test ./...
go run ./cmd/o2c
curl localhost:8080/healthz
```

Configuration is environment-based; see `.env.example`.

## Layout

```
cmd/o2c            service entrypoint
internal/config    env config
internal/health    liveness/readiness
docs               architecture and decisions
deploy             local dependencies (Postgres)
migrations         schema migrations
```
