package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_app/models"
	"log"
	"net/http"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		db: db,
	}
}

func (s *UserService) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = CreateUser(s.db, user.Username, user.FirstName, user.LastName, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func (s *UserService) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("LoginUserHandler called")
}

func CreateUser(db *sql.DB, username string, firstName string, lastName string, password string) error {
	query := "INSERT INTO `users` (`firstName`, `lastName`, `username`, `password`) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, firstName, lastName, username, password)
	if err != nil {
		return err
	}
	return nil
}
