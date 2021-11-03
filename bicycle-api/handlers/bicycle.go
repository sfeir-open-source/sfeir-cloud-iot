package handlers

import (
	"encoding/json"
	"github.com/sfeir-cloud-iot/bicycle-api/models"
	"math"
	"net/http"

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
