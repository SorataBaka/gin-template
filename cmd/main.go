package main

import (
  "fmt"
  "github.com/SorataBaka/gin-template/bootstrap"
  "github.com/SorataBaka/gin-template/api/routes"
)
func main(){
  app := bootstrap.InitializeApp()
  routes.SetupRoute(app)

  address := fmt.Sprintf(":%d", app.Env.PORT)
  app.Engine.Run(address)
}