package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Listen(port string) {
	r := mux.NewRouter()
	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/users", userHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(port, r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte("foo"))
}
