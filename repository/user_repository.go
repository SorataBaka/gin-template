package repository

import (
  "github.com/SorataBaka/gin-template/domain"
	"go.mongodb.org/mongo-driver/mongo"
  "context"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
  database *mongo.Database
  collection string
}

func NewUserRepository(database *mongo.Database, collection string) domain.UserRepository {
  return &userRepository{
    database: database,
    collection: collection,
  }
}

func (repo *userRepository) CreateUser(context context.Context, user *domain.User)(error){
  collection := repo.database.Collection(repo.collection)
  _ , err := collection.InsertOne(context, user)
  if err != nil {
    return err
  }
  return nil
}
func (repo *userRepository)DeleteUserById(context context.Context, ID string)(*domain.User, error){
  collection := repo.database.Collection(repo.collection)
  var user domain.User
  idHex, err := primitive.ObjectIDFromHex(ID)
  if err != nil {
    return &user, err
  }
  err = collection.FindOneAndUpdate(context, bson.M{"_id": idHex}, bson.M{"deleted": true}).Decode(&user)
  return &user ,err
}
func (repo *userRepository)DeleteUserByEmail(context context.Context, email string)(*domain.User, error){
  collection := repo.database.Collection(repo.collection)
  var user domain.User
  err := collection.FindOneAndUpdate(context, bson.M{"email": email}, bson.M{"deleted": true}).Decode(&user)
  return &user, err
}
func (repo *userRepository)GetUserById(context context.Context, id string)(*domain.User, error){
  collection := repo.database.Collection(repo.collection)
  var user domain.User
  idHex, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return &user, err
  } 
  err = collection.FindOne(context, bson.M{"_id": idHex}).Decode(&user)
  return &user, err
}
func (repo *userRepository)GetUserByEmail(context context.Context, email string)(*domain.User, error){
  collection  := repo.database.Collection(repo.collection)
  var user domain.User
  err := collection.FindOne(context, bson.M{"email": email}).Decode(&user)
  if err != nil {
    return nil, err
  }
  return &user, err
}
func (repo *userRepository)ChangePassword(context context.Context, id string, newPassword string)(*domain.User, error){
  collection := repo.database.Collection(repo.collection)
  var user domain.User
  idHex, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return &user, err
  }
  err = collection.FindOneAndUpdate(context, bson.M{"_id": idHex}, bson.M{"password_hash": newPassword}).Decode(&user)
  if err != nil {
    return nil, err
  }
  return &user, err
}
func (repo *userRepository)FetchUsers(context context.Context, filter bson.M)(*[]domain.User, error){
  collection := repo.database.Collection(repo.collection)
  var users []domain.User
  cursor, err := collection.Find(context, filter)
  if err != nil {
    return nil, err
  }
  err = cursor.Decode(&users)
  if err != nil {
    return nil, err
  }
  return &users, err
}