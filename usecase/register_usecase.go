package usecase

import (
	"github.com/SorataBaka/gin-template/bootstrap"
	"github.com/SorataBaka/gin-template/domain"
  "context"
)

type registerUsecase struct {
  app *bootstrap.Application
  repository *domain.UserRepository
}
func NewRegisterUsecase(app *bootstrap.Application, repository *domain.UserRepository) domain.RegisterUsecase{
  return &registerUsecase{
    app: app,
    repository: repository,
  }
}
func (ru *registerUsecase) GetUserByEmail(ginContext context.Context, email string)(*domain.User, error){
  ctx, cancel  := context.WithTimeout(ginContext, *ru.app.TimeoutDuration)
  defer cancel()
  return (*ru.repository).GetUserByEmail(ctx, email)
}
func (ru *registerUsecase) CreateUser(ginContext context.Context, newUser *domain.User)(error){
  ctx, cancel := context.WithTimeout(ginContext, *ru.app.TimeoutDuration)
  defer cancel()
  return (*ru.repository).CreateUser(ctx, newUser)
}