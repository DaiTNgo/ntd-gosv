package handlers

import (
	"context"
	"time"

	"example.com/ntdai/internal/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{
  ID primitive.ObjectID `json:"_id" bson:"_id"`
  CreatedAt time.Time `json:"createdAt" bson:"created_at"`
  UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
  Title string `json:"title" bson:"title"`
}


func CreateProduct(c *fiber.Ctx) error{
  product:= Product{
    ID: primitive.NewObjectID(),
    CreatedAt:time.Now(),
    UpdatedAt:time.Now(),
  }
  if err := c.BodyParser(&product); err != nil{
    return err
  }
    // could also be broken down to be
  // err := c.BodyParser(&product)
  //
  // if err != nil{
  // return err
  // }

  client, err := db.GetMongoClient();

  if err!= nil{
    return err
  }
  collection := client.Database(db.Database).Collection(string(db.ProductsCollection))
  _,err = collection.InsertOne(context.TODO(),product)
}
func GetAllProducts(c *fiber.Ctx){}
