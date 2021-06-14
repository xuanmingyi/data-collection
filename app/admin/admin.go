package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/service", ServiceHandler)

	spa := SPAHandler{StaticPath: "web/dist", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8081", r))
}
