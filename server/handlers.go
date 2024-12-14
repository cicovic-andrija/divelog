package server

import (
	"io/fs"
	"net/http"

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

func multiplexer(includeLocalAPI bool) http.Handler {
	mux := http.NewServeMux()

	// serves as a fallback (frontend)
	mux.HandleFunc("/", indexHandler)
	trace(_https, "handler registered for /")

	// static files (frontend)
	staticFS, _ := fs.Sub(webui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)
	trace(_https, "handler registered for /static/")

	if includeLocalAPI {
		mux.HandleFunc("/data/0", allDataHandler)
		trace(_https, "handler registered for /data/0")

		mux.HandleFunc("/api/fail", func(w http.ResponseWriter, r *http.Request) {
			assert(false, "forced failure")
		})
		trace(_https, "handler registered for /api/fail")
	}

	return mux
}
