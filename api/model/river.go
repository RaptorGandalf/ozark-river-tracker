package model

import "github.com/google/uuid"

type River struct {
	Id        uuid.UUID
	Name      string
	Latitude  float64
	Longitude float64
}
