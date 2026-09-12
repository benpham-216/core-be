# core-be

`core-be` is the Agentic Flow Studio control plane. It owns work-order lifecycle,
human approval gates, policy decisions, and append-only audit events. Provider,
tool, Git, and sandbox execution are adapters and are deliberately outside this
first slice.

## BE-0 / BE-1

- Event-sourced work-order aggregate with explicit lifecycle transitions.
- Optimistic concurrency at the event-store boundary.
- Correlated actor and event metadata.
- HTTP contract for create, read, history, plan approval, and completion.
- PostgreSQL schema for the durable event store and transactional outbox.
- In-memory runtime adapter for deterministic initial tests.

## Lifecycle

`draft -> researching -> solution_review -> planning -> awaiting_plan_approval -> implementing -> validating -> independent_review -> awaiting_publish_approval -> completed`

The aggregate rejects shortcuts, including `implementing -> completed`. Plans and
publishing require a human actor; agents cannot self-approve.

## Local verification

Install Go 1.23 or later, then run:

```sh
go test ./...
go vet ./...
go run ./cmd/api
```

The API uses the in-memory adapter until the PostgreSQL adapter is introduced in
BE-2. Apply `migrations/000001_governance_core.up.sql` before enabling that adapter.
