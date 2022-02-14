package helpers

import (
	"github.com/gorilla/mux"
	"jobExercise/handlers"
)

func CustomRoutes(mux *mux.Router) *mux.Router {
	//USER ROUTES
	mux.HandleFunc("/api/users", handlers.ListUsers).Methods("GET")
	mux.HandleFunc("/api/user/register", handlers.Register).Methods("POST")
	mux.HandleFunc("/api/user", handlers.UpdateUserInfo).Methods("PUT")
	mux.HandleFunc("/api/user", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user", handlers.DeleteUser).Methods("DELETE")

	// VIDEOGAMES ROUTES
	mux.HandleFunc("/api/videogames", handlers.ListVideogames).Methods("POST")
	mux.HandleFunc("/api/videogame/register", handlers.RegisterVideogame).Methods("POST")
	mux.HandleFunc("/api/videogame", handlers.UpdateVideogameInfo).Methods("PUT")
	mux.HandleFunc("/api/videogame", handlers.GetVideogame).Methods("GET")
	mux.HandleFunc("/api/videogame", handlers.DeleteVideogame).Methods("DELETE")

	return mux
}
