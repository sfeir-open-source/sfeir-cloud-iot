package models

type ConfigData struct {
	TargetKm   int     `json:"target_km" example:"150"`
	RadiusInCm float64 `json:"radius_in_cm" example:"30"`
}
