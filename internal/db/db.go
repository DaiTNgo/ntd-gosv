package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

var clientInstance *mongo.Client;
var mongoOnce sync.Once
var clientInstanceError error
type Collection string
const (
  ProductsCollection Collection = "prodcuts"
  // UserCollection Collection = "users"

  )
const (
  URI = "mongodb://localhost:27017"
  Database = "products-api"
  )
func GetMongoClient()(*mongo.Client,error){
  mongoOnce.Do(func(){
    clientOptions :=options.Client().ApplyURI()
    client,err := mongo.Connect(context.TODO(),clientOptions)
    clientInstance = client
    clientInstanceError = err
  })
    return clientInstance, clientInstanceError
}
