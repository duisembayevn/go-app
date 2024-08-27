package users

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go_app/dto"
	"go_app/models"
	"net/http"
	"time"
)

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.store.CreateUser(user.Username, user.FirstName, user.LastName, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func (s *UserService) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body = dto.LoginRequest{}
	json.NewDecoder(r.Body).Decode(&body)

	user, _ := s.store.GetUserByUsername(body.Username)

	if user.Password != body.Password {
		w.WriteHeader(http.StatusUnauthorized)

		var response = dto.ErrorResponse{
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	token, err := GenerateJWT(user.Id, body.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		var response = dto.ErrorResponse{
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	var responseUser = dto.ResponseUser{
		user.Username,
		user.FirstName,
		user.LastName,
	}

	json.NewEncoder(w).Encode(dto.LoginResponse{token, responseUser})
}

func (s *UserService) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	claims := &Claims{}
	token := r.Header.Get("authorization")
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	user, err := s.store.GetUserByUsername(claims.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		var response = dto.ErrorResponse{
			http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized,
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	var response = dto.ResponseUser{
		user.Username,
		user.FirstName,
		user.LastName,
	}

	json.NewEncoder(w).Encode(response)
}

// TODO: refactor jwt stuff

type Claims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`

	jwt.StandardClaims
}

var jwtSecret = []byte("jwt_secret")

func GenerateJWT(userId int, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
