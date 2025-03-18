package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	opentelemetry "github.com/almeida-thiago/city_cep_weather/internals/infra/api/otel"
	usecases "github.com/almeida-thiago/city_cep_weather/internals/use_cases"
	"github.com/almeida-thiago/city_cep_weather/internals/utils"
)

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	span := opentelemetry.InitSpan(r.Context(), "get-cep-weather", "send-to-city-weather")
	defer span.End()

	if r.Method != http.MethodPost {
		RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	type CEPRequest struct {
		CEP string `json:"cep"`
	}

	var req CEPRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error reading body")
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "invalid json")

		return
	}

	cep := req.CEP
	if !utils.IsValidCEP(cep) {
		RespondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	temperature, err := usecases.GetTemperatureFromCep(cep)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}
