package controllers

import (
	"encoding/json"
	"net/http"

	opentelemetry "github.com/almeida-thiago/city_weather/internals/infra/api/otel"
	usecases "github.com/almeida-thiago/city_weather/internals/use_cases"
	utils "github.com/almeida-thiago/city_weather/internals/utils"
)

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	span := opentelemetry.InitSpan(r.Context(), "get-city-weather", "by-zipcode")
	defer span.End()

	cep := r.URL.Query().Get("cep")
	if !utils.IsValidCEP(cep) {
		RespondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	city, err := usecases.GetCityFromCEP(cep)
	if err != nil {
		if err.Error() == "zipcode not found" {
			RespondWithError(w, http.StatusNotFound, "can not find zipcode")
		} else {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	tempC, err := usecases.GetTemperatureForCity(city)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tempResponse := TemperatureResponse{
		TempC: tempC,
		TempF: utils.ConvertCToF(tempC),
		TempK: utils.ConvertCToK(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tempResponse)
}
