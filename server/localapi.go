package server

import (
	"encoding/json"
	"net/http"
)

func allDataHandler(w http.ResponseWriter, r *http.Request) {
	all := &All{
		DiveSites: _inmemDatabase.DiveSites,
		DiveTrips: _inmemDatabase.DiveTrips,
		Dives:     _inmemDatabase.Dives,
	}
	encoded, err := json.Marshal(all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(encoded)
}
