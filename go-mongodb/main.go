package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Student struct {
	Name string
	Age  int
}

var client *mongo.Client

func initDb() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("c: %v \n", client)
	}

	err2 := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err2)
	} else {
		fmt.Println("连接成功！")
	}

}

func insert() {
	s1 := Student{
		Name: "gophist",
		Age:  12,
	}

	c := client.Database("go_db").Collection("Student")
	ior, err := c.InsertOne(context.TODO(), s1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("ior.InsertID: %v \n", ior.InsertedID)
	}

	//defer collection.close
}

func insertMany() {
	c := client.Database("go_db").Collection("Student")
	s := Student{Name: "bingo", Age: 20}
	s1 := Student{Name: "gophist", Age: 20}
	s2 := Student{Name: "bfy", Age: 22}
	students := []interface{}{s, s1, s2}

	imr, err := c.InsertMany(context.TODO(), students)
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}
	fmt.Printf("imr.InsertIDs: %v\n", imr.InsertedIDs)
}

func find() {
	ctx := context.TODO()
	defer client.Disconnect(context.TODO())
	c := client.Database("go_db").Collection("Student")

	c2, err := c.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer c2.Close(ctx)

	for c2.Next(ctx) {
		var result bson.D
		c2.Decode(&result)

		fmt.Printf("result: %v\n", result)
		fmt.Printf("result: %v\n", result.Map())
	}

}

func update() {
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()
	init := bson.D{{"name", "bfy"}}
	update := bson.D{{"$set", bson.D{{"name", "big bfy"}, {"Age", 22}}}}
	rst, err := c.UpdateMany(ctx, init, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result ModifiedCount: %v\n", rst.ModifiedCount)
}

func delete() {
	c := client.Database("go_db").Collection("Student")
	ctx := context.TODO()

	dr, err := c.DeleteMany(ctx, bson.D{{"age", 20}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount)
}

func main() {
	initDb()
	//insert()
	//insertMany()
	//find()
	//update()
	delete()
}

// fm func main
