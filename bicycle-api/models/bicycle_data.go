package models

import (
	"time"
)

type BicycleData struct {
	Revolutions int       `json:"revolutions" example:"10"`
	Rpm         int       `json:"rpm" example:"200"`
	Time        time.Time `json:"time" example:"2021-10-26T22:17:36Z"`
}
