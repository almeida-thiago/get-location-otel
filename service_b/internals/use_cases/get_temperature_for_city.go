package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/almeida-thiago/city_weather/configs"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureForCity(city string) (float64, error) {
	if configs.WEATHER_API_KEY == "" {
		return 0, errors.New("weather API key not set")
	}
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", configs.WEATHER_API_BASE_URL, configs.WEATHER_API_KEY, city)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("failed to get weather data")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var weatherResp WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return 0, err
	}
	return weatherResp.Current.TempC, nil
}
