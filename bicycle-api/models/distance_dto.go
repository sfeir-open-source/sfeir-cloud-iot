package models

import "time"

type DistanceDTO struct {
	Distance float64   `json:"distance" example:"19.5"`
	Start    time.Time `json:"start" example:"1977-04-22T06:00:00Z"`
	End      time.Time `json:"end" example:"2022-01-01T00:00:00Z"`
}
