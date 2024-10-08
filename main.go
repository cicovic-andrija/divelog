package main

import (
	"io/fs"
	"net/http"

	"src.acicovic.me/divelog/webui"
)

func main() {
	srv := &http.Server{
		Addr:    ":8888",
		Handler: router(),
	}

	srv.ListenAndServe()
}

func router() http.Handler {
	mux := http.NewServeMux()

	// index page
	mux.HandleFunc("/", indexHandler)

	// static files
	staticFS, _ := fs.Sub(webui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// api
	mux.HandleFunc("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, there!"))
	})
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := webui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}
	rawFile, _ := webui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}
