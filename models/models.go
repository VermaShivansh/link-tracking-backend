package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	Total_Visits int64
	Countries    map[string]uint64
	Time         []time.Time
}

type Settings struct {
	Accessible_Ips struct {
		Active bool
		Value  []string
	}
	Accessible_Countries struct {
		Active bool
		Value  []string
	}
	Expiry_Time struct {
		Active bool
		Value  time.Time
	}
	Visits_Remaining struct {
		Active bool
		Value  int64
	}
	Scheduled_Usage struct {
		Active     bool
		Start_time time.Time
		End_time   time.Time
	}
}

type LinkUnit struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Shortened_url string
	Target_url    string
	Tagname       string
	Settings      Settings
	Analytics     Analytics
}

// Initialize DS like array, map or so on because they become null as default value rest bool, string, int get default values
func (lu *LinkUnit) Initialize() {
	lu.Id = primitive.NewObjectID()

	// Settings
	lu.Settings.Accessible_Ips.Value = []string{}
	lu.Settings.Accessible_Countries.Value = []string{}

	// Analytics
	lu.Analytics.Countries = map[string]uint64{}
	lu.Analytics.Time = []time.Time{}
}
