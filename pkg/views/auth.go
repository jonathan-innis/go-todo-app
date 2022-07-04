package views

import "github.com/jonathan-innis/go-todo-app/pkg/models"

type RegisterRequest struct {
	Password string `json:"password" validate:"required"`
	User
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId    string `json:"userId"`
	Token     string `json:"token"`
	TokenType string `json:"tokenType"`
}

func NewUserFromRegisterRequest(r *RegisterRequest) *models.User {
	return &models.User{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Metadata: *NewMetadataModel(&r.Metadata),
	}
}
