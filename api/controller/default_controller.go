package controller
import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/SorataBaka/gin-template/domain"

type DefaultController struct {
}
func (defaultController *DefaultController)Default(ginContext *gin.Context){
  ginContext.JSON(http.StatusOK, domain.DefaultResponse{
    Status: 200,
    Message: "OK",
    Code: "OK",
    Path: ginContext.FullPath(),
  })
  
}