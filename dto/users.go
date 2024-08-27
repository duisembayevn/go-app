package dto

// COMMON

type ResponseUser struct {
	Username  string `json:"username"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

// REGISTER

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RegisterResponse struct {
	JwtToken string       `json:"jwtToken"`
	User     ResponseUser `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// LOGIN

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JwtToken string       `json:"jwt_token"`
	User     ResponseUser `json:"user"`
}
