package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/sfeir-cloud-iot/bicycle-api/models"

	"github.com/sfeir-cloud-iot/bicycle-api/db"
)

// GetAllBicyleData return all data from firebase realtime DB
// @Summary Gets all bicycle data from database
// @Description Gets all the bicycle data from database
// @Tags BicyleData
// @Produce json
// @Success 200 {object} []models.BicycleData BicycleData
// @Failure 400 {object} models.ErrorResponse
// @Router /bicycle [get]
func GetAllBicyleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := db.GetAllBicycleData()
	if SendDbErrorIfPresent("bicycle", w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// GetBicycleCurrentSpeed return the current speed of the bicycle based on firebase data
// @Summary Computes the current bicycle speed from database
// @Description Computes the current bicycle speed from the latest data in firebase
// @Tags BicyleSpeed
// @Success 200 {int} current speed of the bicycle
// @Failure 400 {object} models.ErrorResponse
// @Router /bicycle/v1/speed [get]
func GetBicycleCurrentSpeed(w http.ResponseWriter, r *http.Request) {
	lastBicycleData, err := db.GetLastBicycleData()
	if SendDbErrorIfPresent("bicycle", w, err) {
		return
	}

	config, err := db.GetCurrentConfig()
	if SendDbErrorIfPresent("config", w, err) {
		return
	}

	radiusInMeters := config.RadiusInCm / 100
	currentSpeed := 0.12 * math.Pi * radiusInMeters * float64(lastBicycleData.Rpm)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SpeedDTO{Speed: currentSpeed})
}

func GetDistances(w http.ResponseWriter, r *http.Request) {
	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	bicycleData, err := db.GetBicycleDataBetweenDate(start, end)
	if SendDbErrorIfPresent("bicycle", w, err) {
		return
	}

	config, err := db.GetCurrentConfig()
	if SendDbErrorIfPresent("config", w, err) {
		return
	}

	radiusInMeters := config.RadiusInCm / 100
	perimeterWheel := 2 * math.Pi * radiusInMeters

	var start_date = time.Now()
	var end_date = start_date
	distanceInMeter := 0.0
	for _, data := range bicycleData {
		distanceInMeter += perimeterWheel * float64(data.Revolutions)
		if data.Time.Before(start_date) {
			start_date = data.Time
		}
		if data.Time.After(end_date) {
			end_date = data.Time
		}
	}
	distanceTotal := distanceInMeter / 1000
	distance := models.DistanceDTO{Distance: float64(distanceTotal),
		Start: start_date, End: end_date}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(distance)
}
