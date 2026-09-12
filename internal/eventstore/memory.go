package eventstore

import (
	"context"
	"sync"

	"github.com/benpham-216/core-be/internal/domain"
)

type MemoryStore struct { mu sync.RWMutex; events map[string][]domain.Event }

func NewMemoryStore() *MemoryStore { return &MemoryStore{events: make(map[string][]domain.Event)} }

func (s *MemoryStore) Append(_ context.Context, aggregateID string, expectedVersion int, events ...domain.Event) error {
	s.mu.Lock(); defer s.mu.Unlock()
	current := s.events[aggregateID]
	if len(current) != expectedVersion { return ErrConcurrency }
	for i := range events { events[i].Version = expectedVersion+i+1; current = append(current, events[i]) }
	s.events[aggregateID] = current
	return nil
}

func (s *MemoryStore) Load(_ context.Context, aggregateID string) ([]domain.Event, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	return append([]domain.Event(nil), s.events[aggregateID]...), nil
}
