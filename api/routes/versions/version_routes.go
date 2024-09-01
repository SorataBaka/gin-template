package versions

import (
	"github.com/SorataBaka/gin-template/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/SorataBaka/gin-template/api/controller"
	"github.com/SorataBaka/gin-template/api/routes/versions/v1"
)
func Setup(app *bootstrap.Application, baseGroup *gin.RouterGroup){
  versionGroup := baseGroup.Group("/v1")

	defaultController := controller.DefaultController{}
  versionGroup.Any("/", defaultController.Default)

	v1.SetupAuthRoutes(app, versionGroup)
}