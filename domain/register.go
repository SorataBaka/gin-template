package domain
import "context"

type RegisterRequest struct {
  Name string `json:"name" binding:"required"`
  Email string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
  AccessToken string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
}

type RegisterUsecase interface {
  GetUserByEmail(ginContext context.Context, email string)(*User, error)
  CreateUser(ginContext context.Context, newUser *User)(error)
}