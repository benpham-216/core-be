package eventstore

import (
	"context"
	"errors"

	"github.com/benpham-216/core-be/internal/domain"
)

var ErrConcurrency = errors.New("event store concurrency conflict")

type Store interface {
	Append(ctx context.Context, aggregateID string, expectedVersion int, events ...domain.Event) error
	Load(ctx context.Context, aggregateID string) ([]domain.Event, error)
}
