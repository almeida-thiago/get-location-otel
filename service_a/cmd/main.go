package main

import (
	"context"
	"log"

	api "github.com/almeida-thiago/city_cep_weather/internals/infra/api"
	opentelemetry "github.com/almeida-thiago/city_cep_weather/internals/infra/api/otel"
)

func main() {
	otelProvider, err := opentelemetry.InitTracer()
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() { _ = otelProvider.Shutdown(context.Background()) }()
	api.Webserver()
}
