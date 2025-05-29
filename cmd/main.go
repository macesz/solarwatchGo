package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	swHandler "github.com/macesz/solarwatchGo/internal/handler/solarwatch"
	"github.com/macesz/solarwatchGo/internal/service/geolocation"
	"github.com/macesz/solarwatchGo/internal/service/solarwatch"
)


func main() {
	apiKey := os.Getenv("APIKEY");

	print(apiKey)

	geoLocService := geolocation.NewService(apiKey)

	solWatchService := solarwatch.NewService()

	handler := swHandler.NewHandler(geoLocService, solWatchService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api", handler.Handle)

	http.ListenAndServe(":3000", r)
}
