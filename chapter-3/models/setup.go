// package models

// import (
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
// 	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/test2"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.AutoMigrate(&Product{})
// 	DB = db
// }

// uji 1

// package models

// import (
// 	"context"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var DB *mongo.Database

// func ConnectDatabase() {
// 	clientOptions := options.Client().ApplyURI("mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/test")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Connect(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DB = client.Database("test")
// }

// uji 2

package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/test")
	client, err := mongo.NewClient(clientOptions)
	// collection := DB.Collection("test")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("test")
}
