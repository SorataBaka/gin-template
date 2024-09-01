package bootstrap

import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "context"
import "log"
import "time"

func InitializeDatabase(env *EnvObject)(*mongo.Client){
  timeout, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()

  client, err := mongo.Connect(timeout, options.Client().ApplyURI(env.MONGODB_URI))
  if err != nil {
    log.Println(err.Error())
    log.Fatal("Unable to connect to MongoDB Database")
  }
  err = client.Ping(timeout, nil)
  if err != nil {
    log.Fatal("Unable to ping the database")
  }
  return client
}
func CloseDBConnection(client *mongo.Client){
  if client == nil {
    return
  }
  err := client.Disconnect(context.TODO())
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Connection to MongoDB closed")
}