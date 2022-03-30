package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	go handler.pool.Start()

	r := mux.NewRouter()
	r.Use(handler.loggerMiddleware)
	admins := r.PathPrefix("/admins").Subrouter()
	admins.HandleFunc("/artShows", handler.GetArtShowsForAdmins).Methods("GET")
	admins.HandleFunc("/toggleShow/{id}", handler.ToggleShow).Methods("GET")

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/artShows", handler.GetArtShowsForNormalUsers).Methods("GET")
	users.HandleFunc("/getShow/{id}", handler.GetArtShow).Methods("GET")
	users.HandleFunc("/toggleLike/{id}", handler.ToggleLike).Methods("GET")

	socket := r.PathPrefix("/socket").Subrouter()
	socket.HandleFunc("/connect/{id}", handler.ConnectUserToRealTimeUpdates).Methods("GET")

	http.ListenAndServe(":8080", r)
}
