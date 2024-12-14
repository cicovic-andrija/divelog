package server

import (
	"encoding/json"
	"net/http"
)

// Local API; registered only in "dev" mode; error reporting through HTTPS responses is acceptable.

func fetchAll(w http.ResponseWriter, r *http.Request) {
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
	writeJSON(w, encoded)
}

func forceFailure(w http.ResponseWriter, r *http.Request) {
	assert(false, "forced failure")
}
