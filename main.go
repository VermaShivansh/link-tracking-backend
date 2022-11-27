package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ShivanshVerma-coder/link-tracking/db"
	"github.com/ShivanshVerma-coder/link-tracking/helpers"
	"github.com/ShivanshVerma-coder/link-tracking/routes"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	//Load Environments
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error in loading envs")
	}

	// Connect database and load collections
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Unable to connect to db %v", err.Error())
	}

}

func main() {
	// defere function to disconnect databases

	//Check Database connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// Setup routes and run server
	err := routes.SetupRouter().Run()

	if err != nil {
		helpers.PrettyPrint(err)
		log.Fatal("Server crashed")
	}

}
