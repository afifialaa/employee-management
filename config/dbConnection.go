package config

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/afifialaa/employee-management/secrets"

	"context"
	"fmt"
	"log"
)

var UserCollection *mongo.Collection
var EmployeeCollection *mongo.Collection

func DBConnect() {
	clientOptions := options.Client().ApplyURI(secrets.MongoCloud())

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

	// Set database and collection
	UserCollection = client.Database("private").Collection("users")
	EmployeeCollection = client.Database("private").Collection("users")

	// Create index
	_, err = EmployeeCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"email": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		fmt.Println("Email field is not unique")
		log.Fatal(err)
	}
}

// Helper function, duplicate keys
// Looping over the WriteErrors to find code 11000
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
