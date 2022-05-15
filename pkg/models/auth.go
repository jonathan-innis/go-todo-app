package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
