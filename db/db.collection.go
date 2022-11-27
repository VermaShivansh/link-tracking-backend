package db

import "go.mongodb.org/mongo-driver/mongo"

type collections struct {
	LinkStore *mongo.Collection
}

var Collections collections

func loadCollections(client *mongo.Client) {
	Collections.LinkStore = client.Database("link-tracking").Collection("LinkStore")
}
