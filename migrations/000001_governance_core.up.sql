CREATE TABLE work_orders (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  current_version INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE aggregate_events (
  event_id UUID PRIMARY KEY,
  aggregate_id UUID NOT NULL REFERENCES work_orders(id),
  aggregate_type TEXT NOT NULL,
  version INTEGER NOT NULL,
  event_type TEXT NOT NULL,
  actor_type TEXT NOT NULL,
  actor_id TEXT NOT NULL,
  correlation_id UUID NOT NULL,
  payload JSONB NOT NULL,
  occurred_at TIMESTAMPTZ NOT NULL,
  UNIQUE (aggregate_id, version)
);
CREATE INDEX aggregate_events_aggregate_order ON aggregate_events (aggregate_id, version);
CREATE INDEX aggregate_events_correlation ON aggregate_events (correlation_id);

CREATE TABLE outbox_messages (
  id UUID PRIMARY KEY,
  event_id UUID NOT NULL REFERENCES aggregate_events(event_id),
  topic TEXT NOT NULL,
  payload JSONB NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  published_at TIMESTAMPTZ NULL
);
