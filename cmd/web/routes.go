package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/ascii-art", app.ascii)

	return secureHeaders(mux)
}
