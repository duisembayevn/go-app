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

func SetupRoutes(database *sql.DB) *mux.Router {
	r := mux.NewRouter()

	store := db.NewStore(database)

	userService := users.NewUserService(store)
	projectsService := projects.NewProjectsService(store)

	// Routes

	r.HandleFunc("/users/register", userService.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/login", userService.LoginUserHandler).Methods("POST")

	// Routes with JWT middleware
	r.Handle("/users/me", JWTMiddleware(http.HandlerFunc(userService.GetProfileHandler))).Methods("GET")

	r.Handle("/projects/create", JWTMiddleware(http.HandlerFunc(projectsService.CreateProjectHandler))).Methods("POST")
	r.Handle("/projects/{id}", JWTMiddleware(http.HandlerFunc(projectsService.GetProjectHandler))).Methods("GET")
	r.Handle("/projects/{id}", JWTMiddleware(http.HandlerFunc(projectsService.DeleteProjectHandler))).Methods("DELETE")

	r.Handle("/projects", JWTMiddleware(http.HandlerFunc(projectsService.GetProjectsByUserHandler))).Methods("GET")

	return r
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := users.ValidJWT(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
