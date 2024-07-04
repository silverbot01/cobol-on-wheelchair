package main

import (
	"net/http"
	"net/http/cgi"
)

func cgiHandler(w http.ResponseWriter, r *http.Request) {
	handler := cgi.Handler{Path: "the.cow"}
	handler.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", cgiHandler)

	http.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe("localhost:8080", nil)
}
