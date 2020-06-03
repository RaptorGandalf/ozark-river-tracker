package model

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Id           uuid.UUID
	GaugeId      uuid.UUID
	Type         string
	Value        float64
	RecordedDate time.Time //TODO: make this value a datetime in postgres db, currently only recording date
}
