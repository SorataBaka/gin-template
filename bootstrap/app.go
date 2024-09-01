package bootstrap

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
  "go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
  Env *EnvObject
  TimeoutDuration *time.Duration
  Context *context.Context
  Engine *gin.Engine
  Database *mongo.Database
}

func InitializeApp(env *EnvObject, db *mongo.Database)(*Application){
  TimeoutDuration := time.Duration(time.Second * time.Duration(env.DURATION))
  ctx, cancel := context.WithTimeout(context.Background(), TimeoutDuration)
  defer cancel()

  if env.GIN_MODE == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
  engine := gin.Default()
  
  app := &Application{}
  app.Env = env
  app.TimeoutDuration = &TimeoutDuration
  app.Context = &ctx
  app.Engine = engine
  app.Database = db

  return app
}