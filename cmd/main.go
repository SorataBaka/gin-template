package main

import (
  "fmt"
  "github.com/SorataBaka/gin-template/bootstrap"
  "github.com/SorataBaka/gin-template/api/routes"
  "log"
)
func main(){
  env := bootstrap.InitializeEnv()
  db := bootstrap.InitializeDatabase(env)
  app := bootstrap.InitializeApp(env, db.Database(env.MONGODB_DATABASE))

  defer bootstrap.CloseDBConnection(db)

  routes.SetupRoute(app)

  //Start the server
  address := fmt.Sprintf(":%d", app.Env.PORT)
  app.Engine.Run(address)
  log.Println("Server started")
}