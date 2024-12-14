package server

import (
	"net/http"
)

type All struct {
	DiveSites []*DiveSite `json:"dive_sites"`
	DiveTrips []*DiveTrip `json:"dive_trips"`
	Dives     []*Dive     `json:"dives"`
}

func writeJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
