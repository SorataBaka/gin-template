package domain

import "context"
import "go.mongodb.org/mongo-driver/bson/primitive"
import "go.mongodb.org/mongo-driver/bson"


const CollectionUser = "users"

type User struct {
  ID primitive.ObjectID `bson:"_id"`
  Name string `bson:"name"`
  Email string `bson:"email"`
  PasswordHash string `bson:"password_hash"`
  CreatedAt primitive.DateTime `bson:"created_at"`
  Deleted bool `bson:"deleted"`
}

type UserRepository interface {
  CreateUser(context context.Context, user *User) (error)
  DeleteUserById(context context.Context, ID string) (*User, error)
  DeleteUserByEmail(context context.Context, email string) (*User, error)
  GetUserById(context context.Context, id string) (*User, error)
  GetUserByEmail(context context.Context, email string) (*User, error)
  ChangePassword(context context.Context, id string, newPassword string) (*User, error)
  FetchUsers(context context.Context, filter bson.M)(*[]User, error)
}