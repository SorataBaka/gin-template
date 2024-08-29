package controller
import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/SorataBaka/gin-template/domain"

type NotFoundController struct {
}
func (defaultController *NotFoundController)NotFound(ginContext *gin.Context){
  ginContext.JSON(http.StatusNotFound, domain.NotFoundResponse{
    Status: 404,
    Message: "Not Found",
    Code: "NOTFOUND",
  })
}