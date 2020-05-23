package model

import "github.com/google/uuid"

type Gauge struct {
	Id        uuid.UUID
	Name      string
	Code      string
	RiverId   uuid.UUID
	Latitude  float64
	Longitude float64
}
