package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	user2 "go_app/old/service/users"
	"log"
	"net/http"
)

type APIServer struct {
	address  string
	database *sql.DB
}

func NewAPIServer(address string, database *sql.DB) *APIServer {
	return &APIServer{
		address:  address,
		database: database,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user2.NewStore(s.database)
	userHandler := user2.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.address)

	return http.ListenAndServe(s.address, router)
}
