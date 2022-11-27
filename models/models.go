package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	Total_Visits int64             `json:"total_visits"`
	Countries    map[string]uint64 `json:"countries"`
	Time         []time.Time       `json:"time"`
}

type Settings struct {
	Accessible_Ips struct {
		Active bool     `json:"active"`
		Value  []string `json:"value"`
	} `json:"accessible_ips"`
	Accessible_Countries struct {
		Active bool     `json:"active"`
		Value  []string `json:"value"`
	} `json:"accessible_countries"`
	Expiry_Time struct {
		Active bool      `json:"active"`
		Value  time.Time `json:"value"`
	} `json:"expiry_time"`
	Visits_Remaining struct {
		Active bool  `json:"active"`
		Value  int64 `json:"value"`
	} `json:"visits_remaining"`
	Scheduled_Usage struct {
		Active     bool      `json:"active"`
		Start_time time.Time `json:"start_time"`
		End_time   time.Time `json:"end_time"`
	} `json:"scheduled_usage"`
}

type LinkUnit struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Shortened_url string             `json:"shortened_url"`
	Target_url    string             `json:"target_url"`
	Tagname       string             `json:"tagname"`
	Settings      Settings           `json:"settings"`
	Analytics     Analytics          `json:"analytics"`
}

// Initialize DS like array, map or so on because they become null as default value rest bool, string, int get default values
func (lu *LinkUnit) Initialize() {
	lu.Id = primitive.NewObjectID()

	// Settings
	lu.Settings.Accessible_Ips.Value = []string{}
	lu.Settings.Accessible_Countries.Value = []string{}
	lu.Settings.Expiry_Time.Value = time.Now()

	// Analytics
	lu.Analytics.Countries = map[string]uint64{}
	lu.Analytics.Time = []time.Time{}
}
