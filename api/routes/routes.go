package routes

import (
	"github.com/SorataBaka/gin-template/api/controller"
	"github.com/SorataBaka/gin-template/bootstrap"
	"github.com/SorataBaka/gin-template/api/routes/versions"
)
func SetupRoute(app *bootstrap.Application){
  apiRouter := app.Engine.Group("/api")
  defaultController := controller.DefaultController{}
  apiRouter.Any("/", defaultController.Default)
  
  versions.Setup(app, apiRouter)

  noRouteController := controller.NotFoundController{}
  app.Engine.NoRoute(noRouteController.NotFound)

}