package v1

import "github.com/gin-gonic/gin"
import "github.com/SorataBaka/gin-template/bootstrap"
import "github.com/SorataBaka/gin-template/api/controller"
import "github.com/SorataBaka/gin-template/api/controller/auth"
import "github.com/SorataBaka/gin-template/usecase"
import "github.com/SorataBaka/gin-template/repository"

func SetupAuthRoutes(app *bootstrap.Application, baseGroup *gin.RouterGroup){
  authGroup := baseGroup.Group("/auth")

  defaultController := controller.DefaultController{}
  authGroup.Any("/", defaultController.Default)
  userRepository := repository.NewUserRepository(app.Database, "users")


  registerController := auth.RegisterController{
    App: app,
    Usecase: usecase.NewRegisterUsecase(app, &userRepository),
  }
  authGroup.POST("/register", registerController.Register)
}