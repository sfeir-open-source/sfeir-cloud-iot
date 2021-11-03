package handlers

import (
	"github.com/sfeir-cloud-iot/bicycle-api/models"
	"net/http"
)

func SendDbErrorIfPresent(domain string, w http.ResponseWriter, err error) bool {
	if err != nil {
		var errResponse models.UnprocessableEntity
		errorDTO := errResponse.GetError("Couldn't read " + domain + " data : " + err.Error())
		http.Error(w, errorDTO.Message, errorDTO.StatusCode)
		return true
	}
	return false
}
