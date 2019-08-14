package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id 		primitive.ObjectID 	`bson:"_id"`
	Name 	string	`bson:"name"`
	Sex		string	`bson:"sex"`
	Age 	int 	`bson:"age"`
}

func main()  {
	mongodbUrl := "mongodb://127.0.0.1:27017"
	//mongodbUrl := "mongodb://admin:admin@127.0.0.1:27017/table"
	clientOptions := options.Client().ApplyURI(mongodbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed")
		return
	}

	// 连接集合
	collection := client.Database("table").Collection("test")


	// 添加/插入

	// 添加单个
	if insertId, err := collection.InsertOne(context.TODO(), User{Name: "Nick"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(insertId)	// &{ObjectID("5d5391f7b9023b91ee179295")}
		fmt.Println("Insert successfully.")
	}

	// 添加多个
	users := []interface{}{
		User{Name:"Mark", Age:18, Sex:"male"},
		User{Name:"Ark", Age:12},
		User{Name:"Bob", Sex:"male"},
	}
	if insertId, err := collection.InsertMany(context.TODO(), users); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(insertId)
		// 输出结果：&{[ObjectID("5d5391f7b9023b91ee179296") ObjectID("5d5391f7b9023b91ee179297") ObjectID("5d5391f7b9023b91ee179298")]}
		fmt.Println("Insert successfully.")
	}


	// 查询

	// 获得单个结果
	user := User{}
	if err := collection.FindOne(context.TODO(), bson.M{"age": 18}).Decode(&user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)	// {Mark male 18}
		fmt.Println("Find successfully.")
	}

	// 获得多个结果
	cur, err := collection.Find(context.TODO(), bson.M{"sex": "male"})

	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var user User
		if err := cur.Decode(&user); err != nil {
			fmt.Println(err)
		}
		fmt.Println(user.Name, user.Age)
	}
	fmt.Println("Find successfully.")


	// 修改

	// 更新
	if modifyResult, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"name": "Mark"},
		bson.M{"$set": User{Name:"Lily", Sex:"female"}}); err != nil {
			fmt.Println(err)
	} else {
		fmt.Println(modifyResult)	// &{1 1 0 <nil>}
		fmt.Println("Update successfully.")
	}

	// 替换
	if modifyResult, err := collection.ReplaceOne(
		context.TODO(),
		bson.M{"name": "Nick"},
		User{Name: "Jack", Age: 25, Sex: "female"}); err != nil {
			fmt.Println(err)
	} else {
		fmt.Println(modifyResult)	// &{1 1 0 <nil>}
		fmt.Println("Replace successfully.")
	}


	// 删除

	if deleteCount, err := collection.DeleteOne(context.TODO(), bson.M{"age": "Nick"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(deleteCount)	// &{0}
		fmt.Println("Delete successfully.")
	}

	// 删除多个
	if deleteCount, err := collection.DeleteMany(context.TODO(), bson.M{"sex": "male"}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(deleteCount)	// &{2}
		fmt.Println("Delete successfully.")
	}


	// 清空集合
	if _, err := collection.DeleteMany(context.TODO(), bson.D{}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Clear")
}