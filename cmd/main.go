package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	swHandler "github.com/macesz/solarwatchGo/internal/handler/solarwatch"
	"github.com/macesz/solarwatchGo/internal/service/geolocation"
	"github.com/macesz/solarwatchGo/internal/service/solarwatch"
)

const apiKey = "3c4308a7a3e66d61d8e7f4b1cc5ec4bc"

func main() {
	geoLocService := geolocation.NewService(apiKey)

	solWatchSerice := solarwatch.NewService()

	handler := swHandler.NewHandler(geoLocService, solWatchSerice)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api", handler.Handle)

	http.ListenAndServe(":3000", r)
}
