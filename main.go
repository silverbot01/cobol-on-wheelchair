package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func cgiHandler(w http.ResponseWriter, r *http.Request) {
	handler := cgi.Handler{Path: "the.exe"}

	var tabstring []byte = make([]byte, 0, 30)
	diff := 30 - len(r.RequestURI)
	for i := 0; i < diff; i++ {
		tabstring = append(tabstring, ' ')
	}
	log.Println("\t:: " + r.RemoteAddr + "\t:: " + r.Method + "\t:: " + r.RequestURI + string(tabstring) + ":: FROM :: " + r.Host)

	handler.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", cgiHandler)

	http.Handle("/static/*", loggerMW(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}

func loggerMW(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tabstring []byte = make([]byte, 0, 30)
		diff := 30 - len(r.RequestURI)
		for i := 0; i < diff; i++ {
			tabstring = append(tabstring, ' ')
		}
		log.Println("\t:: " + r.RemoteAddr + "\t:: " + r.Method + "\t:: " + r.RequestURI + string(tabstring) + ":: FROM :: " + r.Host)
		h.ServeHTTP(w, r)
	})
}
