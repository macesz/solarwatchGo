package solarwatch

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	geoLocationGetter GeoLocationGetter
	solarWatchGetter  SolarWatchGetter
}

func NewHandler(
	geoLocationGetter GeoLocationGetter,
	solarWatchGetter SolarWatchGetter,
) *Handler {
	return &Handler{
		geoLocationGetter: geoLocationGetter,
		solarWatchGetter:  solarWatchGetter,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	city := q.Get("city")
	if city == "" {
		http.Error(w, "city parameter is required", http.StatusBadRequest)
		return
	}

	countryCode := q.Get("countryCode")
	if countryCode == "" {
		http.Error(w, "countryCode parameter is required", http.StatusBadRequest)
		return
	}

	state := q.Get("state")

	datestr := q.Get("date")
	if datestr == "" {
		http.Error(w, "date parameter is required", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	geoLoc, err := h.geoLocationGetter.GetReport(city, countryCode, state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sw, err := h.solarWatchGetter.GetReport(geoLoc.Lat, geoLoc.Lon, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &dto{
		Sunset:  sw.Sunset,
		Sunrise: sw.Sunrise,
	}

	body, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(body)

	w.WriteHeader(http.StatusOK)
}
