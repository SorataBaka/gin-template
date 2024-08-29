package versions

import (
	"github.com/SorataBaka/gin-template/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/SorataBaka/gin-template/api/controller"
)
func Setup(app *bootstrap.Application, baseGroup *gin.RouterGroup){
  versionGroup := baseGroup.Group("/v1")

	defaultController := controller.DefaultController{}
  versionGroup.Any("/", defaultController.Default)
}