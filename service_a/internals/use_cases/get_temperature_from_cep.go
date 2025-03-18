package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/almeida-thiago/city_cep_weather/configs"
)

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetTemperatureFromCep(cep string) (temperature TemperatureResponse, err error) {
	resp, err := http.Get(fmt.Sprintf("%s?cep=%s", configs.CITY_WEATHER_API_BASE_URL, cep))
	if err != nil {
		return temperature, errors.New("error while fetching weather data")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return temperature, errors.New("error while reading weather data")
	}

	if err = json.Unmarshal(body, &temperature); err != nil {
		return temperature, errors.New("error while processing weather data")
	}

	return temperature, nil
}
