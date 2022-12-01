package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/ShivanshVerma-coder/link-tracking/db"
	"github.com/ShivanshVerma-coder/link-tracking/helpers"
	"github.com/ShivanshVerma-coder/link-tracking/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateLink(shortened_url_id string, target_url string, tagName string) (string, error) {
	linkUnit := &models.LinkUnit{}
	linkUnit.Initialize()

	linkUnit.Target_url = target_url
	linkUnit.Shortened_url = shortened_url_id
	linkUnit.Tagname = tagName

	insertedResult, err := db.Collections.LinkStore.InsertOne(context.TODO(), linkUnit)

	if err != nil {
		log.Fatal("Error in inserting document")
		return "", err
	}

	helpers.PrettyPrint(insertedResult)

	return shortened_url_id, nil
}

func GetTargetLink(shortened_url_id string) (models.LinkUnit, error) {
	filter := bson.M{
		"shortened_url": shortened_url_id,
	}
	opts := options.FindOne().SetProjection(bson.M{"target_url": 1, "_id": 1, "settings": 1})
	linkUnit := models.LinkUnit{}

	err := db.Collections.LinkStore.FindOne(context.TODO(), filter, opts).Decode(&linkUnit)

	if err != nil {
		fmt.Printf("Error in getting target link %v", err)
		return models.LinkUnit{}, err
	}

	return linkUnit, nil
}

func GetCompleteInfo(shortened_url_id string) (models.LinkUnit, error) {
	filter := bson.M{
		"shortened_url": shortened_url_id,
	}
	linkUnit := models.LinkUnit{}

	err := db.Collections.LinkStore.FindOne(context.TODO(), filter).Decode(&linkUnit)

	if err != nil {
		fmt.Printf("Error in getting target link %v", err)
		return models.LinkUnit{}, err
	}

	return linkUnit, nil
}

func UpdateSettings(shortened_url_id string, newLinkUnit models.LinkUnit) error {
	filter := bson.M{
		"shortened_url": shortened_url_id,
	}
	update := bson.M{
		"$set": bson.M{
			"settings":   newLinkUnit.Settings,
			"target_url": newLinkUnit.Target_url,
		},
	}

	result, err := db.Collections.LinkStore.UpdateOne(context.TODO(), filter, update)
	fmt.Println(result)
	if err != nil {
		fmt.Printf("Error in getting target link %v", err)
		return err
	}

	return nil
}
