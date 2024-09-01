package auth

import "github.com/gin-gonic/gin"
import "github.com/SorataBaka/gin-template/bootstrap"
import "github.com/SorataBaka/gin-template/domain"
import "net/http"
import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterController struct {
  App *bootstrap.Application
  Usecase domain.RegisterUsecase
}

func (controller *RegisterController) Register(ginContext *gin.Context){
  var request domain.RegisterRequest
  err := ginContext.ShouldBindJSON(&request)
  if err != nil {
    ginContext.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{
      Status: 400,
      Message: err.Error(),
      Code: "BADREQUEST",
    })
    return
  }
  _, err = controller.Usecase.GetUserByEmail(ginContext, request.Email)
  if err == nil {
    ginContext.AbortWithStatusJSON(http.StatusBadRequest, domain.ErrorResponse{
      Status: 400,
      Message: "Email is already used",
      Code: "BADREQUEST",
    })
    return
  }
  newUserObject := domain.User{
    Name: request.Name,
    Email: request.Email,
    PasswordHash: request.Password,
    CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
    Deleted: false,
  }
  err = controller.Usecase.CreateUser(ginContext, &newUserObject)
  if err != nil {
    ginContext.AbortWithStatusJSON(http.StatusInternalServerError, domain.ErrorResponse{
      Status: 500,
      Message: err.Error(),
      Code: "INTERNALSERVERERROR",
    })
    return
  }
  ginContext.JSON(http.StatusOK, domain.OKResponse{
    Status: 200,
    Message: "OK",
    Code: "OK",
    Data: request,
  })
}