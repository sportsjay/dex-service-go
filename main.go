package main

import (
	"fmt"
	"net/http"

	h "service/dex/handlers"

	"github.com/gorilla/mux"
	// "github.com/gorilla/websocket"
)

type msg struct {
	Num int
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func main() {
	router := mux.NewRouter()

	// Handle middlewares
	router.Use(defaultHeaderMiddleware)

	// Handle routes

	// Order
	ordersR := router.PathPrefix("/order").Subrouter()
	// healthchecks.Use(MiddlewareOne)
	ordersR.Use(authenticationMiddleware)
	ordersR.Path("/").Methods(http.MethodGet).HandlerFunc(h.HOrderGet)
	ordersR.Path("/").Methods(http.MethodPost).HandlerFunc(h.HOrderPost)

	// protected := router.PathPrefix("/admin").Subrouter()
	// // protected.Use(MiddlewareTwo)
	// protected.HandleFunc("/dashboard", handler)

	port := ":5000"
	fmt.Printf("Server running in port%s", port)
	panic(http.ListenAndServe(port, router))
}

func defaultHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		// r.Header.Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
