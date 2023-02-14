package billing

import (
	"encoding/json"
	"time"
)

type AbsoluteEvent struct {
	IdempotencyKey string    `json:"idempotency_key"`
	MetricName     string    `json:"metric"`
	Type           string    `json:"type"`
	TenantID       string    `json:"tenant_id"`
	TimelineID     string    `json:"timeline_id"`
	Time           time.Time `json:"time"`
	Value          int       `json:"value"`
}

func (e *AbsoluteEvent) MarshalJSON() ([]byte, error) {
	e.Type = "absolute"
	return json.Marshal(e)
}

type IncrementalEvent struct {
	IdempotencyKey string    `json:"idempotency_key"`
	MetricName     string    `json:"metric"`
	Type           string    `json:"type"`
	EndpointID     string    `json:"endpoint_id"`
	StartTime      time.Time `json:"start_time"`
	StopTime       time.Time `json:"stop_time"`
	Value          int       `json:"value"`
}

func (e *IncrementalEvent) MarshalJSON() ([]byte, error) {
	e.Type = "incremental"
	return json.Marshal(e)
}