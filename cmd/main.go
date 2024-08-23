package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go_app/config"
	"go_app/db"
	"go_app/services/projects"
	"go_app/services/users"
	"log"
	"net/http"
)

func main() {
	db := db.ConnectDB(config.Envs)
	r := SetupRoutes(db)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	userService := users.NewUserService(db)
	r.HandleFunc("/users/register", userService.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/login", userService.LoginUserHandler).Methods("POST")

	projectsService := projects.NewProjectsService(db)
	r.HandleFunc("/projects/create", projectsService.CreateProjectHandler).Methods("POST")
	r.HandleFunc("/projects/findById/{id}", projectsService.GetProjectByIdHandler).Methods("GET")
	r.HandleFunc("/projects/findByUser/{id}", projectsService.GetProjectsByUserIdHandler).Methods("GET")

	return r
}
