package config

import (
	mux2 "github.com/gorilla/mux"
	"jobExercise/database"
	"jobExercise/helpers"
	"log"
	"net/http"
)

type Server struct {
	Addr string
	Handler http.Handler
}

func initializeConfig(mux http.Handler) Server{
	serverConfig := Server{
		Addr: ":3000",
		Handler: mux,
	}
	return serverConfig
}

func StartServer(){
	mux := mux2.NewRouter()
	routes := helpers.CustomRoutes(mux)
	serverConfig := initializeConfig(routes)
	server := &http.Server{
		Addr: serverConfig.Addr,
		Handler: serverConfig.Handler,
	}
	database.MongoDBInitialization()
	log.Fatalln(server.ListenAndServe())
}