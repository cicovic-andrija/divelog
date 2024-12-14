package server

import (
	"net/http"
)

type All struct {
	DiveSites []*DiveSite `json:"dive_sites"`
	DiveTrips []*DiveTrip `json:"dive_trips"`
	Dives     []*Dive     `json:"dives"`
}

type SiteHead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func writeJSON(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _serverControl.corsAllowAll {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	w.Write(data)
}
