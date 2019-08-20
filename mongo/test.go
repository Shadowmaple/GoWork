package main

//import (
//	"context"
//	"log"
//	"time"
//
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func Test() {
//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	//if err := client.Connect(ctx); err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//col := client.Database("table").Collection("test")
//	//
//	//if _, err := col.InsertOne(ctx, bson.D{{"name", "Shadow"}}); err != nil {
//	//	log.Fatal(err)
//	//}
//
//}