package db

import (
	"context"
	"fmt"
	"log"
	"prometheus/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect shall connect to MongoDB
func Connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("C").Collection("adheip")

	return collection
}

func CreateNode(url string) {

	values := queries.CPUNode(url)

	b, err := bson.Marshal(values)
	if err != nil {
		log.Fatalf("%s", err)
	}

	collection := Connect()

	result, err := collection.InsertOne(context.TODO(), b)

	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("%s", result.InsertedID)

}
