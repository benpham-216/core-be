package domain

import "time"

type Event struct {
	ID            string         `json:"id"`
	AggregateID   string         `json:"aggregateId"`
	AggregateType string         `json:"aggregateType"`
	Version       int            `json:"version"`
	Type          string         `json:"type"`
	Actor         Actor          `json:"actor"`
	CorrelationID string         `json:"correlationId"`
	OccurredAt    time.Time      `json:"occurredAt"`
	Payload       map[string]any `json:"payload"`
}
