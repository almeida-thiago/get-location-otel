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

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro,omitempty"`
}

func GetCityFromCEP(cep string) (string, error) {
	url := fmt.Sprintf("%s/%s/json/", configs.VIACEP_API_BASE_URL, cep)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var viaCEPResp ViaCEPResponse
	if err := json.Unmarshal(body, &viaCEPResp); err != nil {
		return "", err
	}

	if viaCEPResp.Erro {
		return "", errors.New("zipcode not found")
	}
	if viaCEPResp.Localidade == "" {
		return "", errors.New("city not found in CEP data")
	}
	return viaCEPResp.Localidade, nil
}
