package server

import (
	"encoding/json"
	"io/fs"
	"net/http"
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
		sites := _inmemDatabase.DiveSites[1:]
		resp, err = json.Marshal(sites)
	}

	if err != nil {
		trace(_error, "http: failed to marshal dive site data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSON(w, resp)
}

func fetchSite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	siteID, _ := strconv.Atoi(id)
	site := _inmemDatabase.DiveSites[siteID]
	resp, _ := json.Marshal(site)
	writeJSON(w, resp)
}

func multiplexer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /data/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	trace(_https, "handler registered for /data/")

	mux.HandleFunc("GET /data/sites", fetchSites)
	trace(_https, "handler registered for /data/sites")

	mux.HandleFunc("GET /data/sites/{id}", fetchSite)
	trace(_https, "handler registered for /data/sites/{id}")

	// index / root (frontend)
	mux.HandleFunc("/", indexHandler)
	trace(_https, "handler registered for /")

	// static files (frontend)
	staticFS, _ := fs.Sub(webui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)
	trace(_https, "handler registered for /static/")

	if _serverControl.localAPI {
		mux.HandleFunc("GET /data/0", fetchAll)
		trace(_https, "handler registered for /data/0")

		mux.HandleFunc("/action/fail", forceFailure)
		trace(_https, "handler registered for /action/fail")
	}

	return mux
}
