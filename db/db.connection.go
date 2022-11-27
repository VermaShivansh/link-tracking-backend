package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ip2location/ip2location-go/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var IPClient *ip2location.DB

func ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	MongoClient = client
	fmt.Println("Connected to DB successfully")

	loadCollections(MongoClient)
	fmt.Println("Collections loaded successfully")

	connectIPDB()

	return nil
}

func connectIPDB() {
	// Setup IP
	db, err := ip2location.OpenDB("IP2LOCATION-LITE-DB11.IPV6.BIN")
	if err != nil {
		log.Fatal(err)
	}

	IPClient = db
}
