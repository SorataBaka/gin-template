package bootstrap

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Application struct {
  Env *EnvObject
  TimeoutDuration *time.Duration
  Context *context.Context
  Engine *gin.Engine
}

func InitializeApp()(*Application){
  env := InitializeEnv()
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

  return app
}