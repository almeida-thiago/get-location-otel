package configs

import "os"

const SERVICE_NAME = "city_weather"

const OTEL_COLLECTOR = "localhost:4317"

const VIACEP_API_BASE_URL = "https://viacep.com.br/ws"
const WEATHER_API_BASE_URL = "http://api.weatherapi.com/v1"

var WEATHER_API_KEY = os.Getenv("WEATHER_API_KEY")
