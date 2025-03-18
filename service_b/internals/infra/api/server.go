package api

import (
	"log"
	"net/http"
	"os"

	"github.com/almeida-thiago/city_weather/internals/infra/api/controllers"
)

func Webserver() {
	http.HandleFunc("/weather", controllers.WeatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Servidor iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
