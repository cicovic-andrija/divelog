package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"slices"
	"sort"
	"strconv"

	"src.acicovic.me/divelog/webui"
)

// frontend-related handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := webui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := webui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func fetchSites(w http.ResponseWriter, r *http.Request) {
	var (
		resp []byte
		err  error
	)

	if r.URL.Query().Get("headonly") == "true" {
		heads := make([]*SiteHead, 0, len(_inmemDatabase.DiveSites))
		for _, site := range _inmemDatabase.DiveSites[1:] {
			heads = append(heads, &SiteHead{
				ID:   site.ID,
				Name: site.Name,
			})
		}
		sort.Slice(heads, func(i, j int) bool {
			return heads[i].Name < heads[j].Name
		})
		resp, err = json.Marshal(heads)
	} else {
		sites := []*SiteFull{}
		for _, site := range _inmemDatabase.DiveSites[1:] {
			sites = append(sites, NewSiteFull(site, _inmemDatabase.Dives[1:]))
		}
		resp, err = json.Marshal(sites)
	}

	if err != nil {
		trace(_error, "http: failed to marshal dive site data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func fetchSite(w http.ResponseWriter, r *http.Request) {
	// TODO: validate ID
	id := r.PathValue("id")
	siteID, _ := strconv.Atoi(id)
	site := _inmemDatabase.DiveSites[siteID]

	resp, err := json.Marshal(NewSiteFull(site, _inmemDatabase.Dives[1:]))
	if err != nil {
		trace(_error, "http: failed to marshal single dive site data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func fetchTrips(w http.ResponseWriter, r *http.Request) {
	trips := make([]*Trip, 0, len(_inmemDatabase.DiveTrips))
	reverse := r.URL.Query().Get("reverse") == "true"
	if reverse {
		for _, trip := range _inmemDatabase.DiveTrips[1:] {
			trips = append(trips, &Trip{
				ID:    trip.ID,
				Label: trip.Label,
			})
		}
	} else {
		for i := len(_inmemDatabase.DiveTrips) - 1; i > 0; i-- {
			trips = append(trips, &Trip{
				ID:    _inmemDatabase.DiveTrips[i].ID,
				Label: _inmemDatabase.DiveTrips[i].Label,
			})
		}
	}

	for _, trip := range trips {
		for _, dive := range _inmemDatabase.Dives[1:] {
			if dive.DiveTripID == trip.ID {
				// TODO: assert _inmemDatabase.DiveSites[dive.DiveSiteID]
				trip.LinkedDives = append(trip.LinkedDives, NewDiveHead(dive, _inmemDatabase.DiveSites[dive.DiveSiteID]))
			}
		}
		if !reverse {
			slices.Reverse(trip.LinkedDives)
		}
	}

	resp, err := json.Marshal(trips)
	if err != nil {
		trace(_error, "http: failed to marshal dive trip data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func fetchDives(w http.ResponseWriter, r *http.Request) {
	var (
		resp []byte
		err  error
		tag  = r.URL.Query().Get("tag")
	)

	if r.URL.Query().Get("headonly") == "true" {
		heads := make([]*DiveHead, 0, len(_inmemDatabase.Dives))
		for _, dive := range _inmemDatabase.Dives[1:] {
			// TODO: assert _inmemDatabase.DiveSites[dive.DiveSiteID]
			heads = append(heads, NewDiveHead(dive, _inmemDatabase.DiveSites[dive.DiveSiteID]))
		}
		resp, err = json.Marshal(heads)
	} else {
		dives := []*DiveFull{}
		for _, dive := range _inmemDatabase.Dives[1:] {
			if dive.IsTaggedWith(tag) {
				// TODO: assert _inmemDatabase.DiveSites[dive.DiveSiteID]
				dives = append(dives, NewDiveFull(dive, _inmemDatabase.DiveSites[dive.DiveSiteID]))
			}
		}
		resp, err = json.Marshal(dives)
	}

	if err != nil {
		trace(_error, "http: failed to marshal dive data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func fetchDive(w http.ResponseWriter, r *http.Request) {
	// TODO: validate ID
	id := r.PathValue("id")
	diveID, _ := strconv.Atoi(id)
	dive := _inmemDatabase.Dives[diveID]

	// TODO: assert _inmemDatabase.DiveSites[dive.DiveSiteID]
	resp, err := json.Marshal(NewDiveFull(dive, _inmemDatabase.DiveSites[dive.DiveSiteID]))
	if err != nil {
		trace(_error, "http: failed to marshal single dive data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func fetchTags(w http.ResponseWriter, r *http.Request) {
	tags := make(map[string]int)
	for _, dive := range _inmemDatabase.Dives[1:] {
		for _, tag := range dive.Tags {
			tags[tag]++
		}
	}

	resp, err := json.Marshal(tags)
	if err != nil {
		trace(_error, "http: failed to marshal tags data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	send(w, resp)
}

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	// data handlers
	mux.HandleFunc("GET /data/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	trace(_https, "handler registered for /data/")

	mux.HandleFunc("GET /data/sites", fetchSites)
	trace(_https, "handler registered for /data/sites")
	// TODO: /data/sites/{$} returns 404

	mux.HandleFunc("GET /data/sites/{id}", fetchSite)
	trace(_https, "handler registered for /data/sites/{id}")

	mux.HandleFunc("GET /data/trips", fetchTrips)
	trace(_https, "handler registered for /data/trips")
	// TODO: /data/trips/{$} returns 404

	mux.HandleFunc("GET /data/dives", fetchDives)
	trace(_https, "handler registered for /data/dives")
	// TODO: /data/dives/{$} returns 404

	mux.HandleFunc("GET /data/dives/{id}", fetchDive)
	trace(_https, "handler registered for /data/dives/{id}")

	mux.HandleFunc("GET /data/tags", fetchTags)
	trace(_https, "handler registered for /data/tags")
	// TODO: /data/tags/{$} returns 404

	// index / root (frontend)
	mux.HandleFunc("/", indexHandler)
	trace(_https, "handler registered for /")

	// static files (frontend)
	staticFS, _ := fs.Sub(webui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)
	trace(_https, "handler registered for /static/")

	// local API handlers
	if _serverControl.localAPI {
		mux.HandleFunc("GET /data/0", fetchAll)
		trace(_https, "handler registered for /data/0")

		mux.HandleFunc("POST /action/fail", forceFailure)
		trace(_https, "handler registered for /action/fail")

		mux.HandleFunc("POST /action/rebuild", rebuildDatabase)
		trace(_https, "handler registered for /action/rebuild")
	}

	return mux
}

func send(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _serverControl.corsAllowAll {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	_, err := w.Write(data)
	if err != nil {
		trace(_error, "send: %v", err)
	}
}
