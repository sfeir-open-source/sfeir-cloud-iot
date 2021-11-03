package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sfeir-cloud-iot/bicycle-api/db"
)

// GetConfig return the current config stored in Firebase
// @Summary Gets the current config from database
// @Description Gets the current config from database
// @Tags BicyleConfig
// @Produce json
// @Success 200 {object} models.BicyleConfig BicycleConfig
// @Failure 400 {object} models.ErrorResponse
// @Router /config [get]
func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := db.GetCurrentConfig()
	if SendDbErrorIfPresent("config", w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
