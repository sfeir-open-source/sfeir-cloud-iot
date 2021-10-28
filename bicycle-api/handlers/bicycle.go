package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sfeir-cloud-iot/bicycle-api/db"
)

// GetAllBicyleData return all data from firebase realtime DB
// @Summary Get users from database
// @Description Get users array objects data from database
// @Tags BicyleData
// @Produce json
// @Success 200 {object} []models.BicycleData BicycleData
// @Failure 400 {object} models.ErrorResponse
// @Router /bicycle [get]
func GetAllBicyleData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, _ := db.GetAllData()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
