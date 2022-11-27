package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ShivanshVerma-coder/link-tracking/db"
	"github.com/ShivanshVerma-coder/link-tracking/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAnalytics() (string, error) {
	return "", nil
}

func UpdateAnalytics(linkUnit models.LinkUnit, countryCode string) error {
	update := bson.M{"$inc": bson.M{
		"analytics.total_visits":                           1,
		fmt.Sprintf("analytics.countries.%v", countryCode): 1,
	},
		"$push": bson.M{"analytics.time": time.Now()},
		"$set":  bson.M{"settings.visits_remaining.value": linkUnit.Settings.Visits_Remaining.Value - 1},
	}
	filter := bson.M{"_id": linkUnit.Id}

	_, err := db.Collections.LinkStore.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
