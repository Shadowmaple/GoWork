package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name 	string	`bson:"name"`
}

func main()  {
	// mongodbUrl := "mongodb://127.0.0.1:27017"
	mongodbUrl := "mongodb://admin:admin@127.0.0.1:27017/table"
	clientOptions := options.Client().ApplyURI(mongodbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed")
		return
	}

	collection := client.Database("table").Collection("xk")
	_, err = collection.InsertOne(context.TODO(), User{Name: "Nick"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert Success")
}