# Architecture

## Service shape (v1)

Modular monolith first, extracted into services only when a boundary is proven under load:

```
cmd/o2c            HTTP API + lifecycle
internal/order     order aggregate and state machine
internal/inventory reservation / allocation core
internal/ledger    double-entry postings, costing
internal/settlement netting and receipts
internal/outbox    transactional outbox + dispatcher
```

## Consistency model

| Domain | Guarantee | Mechanism |
|--------|-----------|-----------|
| Reservation per warehouse | strong | row lock on inventory position; conditional decrement |
| Ledger | strong + balanced | single transaction per posting batch; sum(debits) == sum(credits) |
| Cross-warehouse allocation | causal with compensation | reservation hierarchy, explicit release |
| Settlement finality | atomic per netting run | single DB transaction + signed receipt |
| Cross-entity notifications | eventual | transactional outbox |

## Invariants

1. `available = on_hand - reserved - quarantined` for every (tenant, warehouse, sku).
2. No reservation is confirmed unless the decrement succeeded; no negative available.
3. Every inventory movement with financial impact emits balanced journal entries.
4. Every external command carries an idempotency key; replays return the original result.
5. Every state transition is append-only and attributable.

## Open decisions

Current implementation defaults:

- Netting: multilateral, per settlement window.
- Receipt: signed canonical JSON payload (ed25519), transparency log deferred.
- Soft reservations: yes, with expiry.
