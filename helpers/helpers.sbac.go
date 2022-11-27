package helpers

import (
	"fmt"
	"time"

	"github.com/ShivanshVerma-coder/link-tracking/models"
	"github.com/ip2location/ip2location-go/v9"
)

func SBAC(linkUnit models.LinkUnit, ip string, record ip2location.IP2Locationrecord) (bool, error) {
	countryCode := record.Country_short
	fmt.Println(linkUnit)
	if linkUnit.Settings.Accessible_Countries.Active {
		found := false
		for _, val := range linkUnit.Settings.Accessible_Countries.Value {
			if val == countryCode {
				found = true
			}
		}
		if !found {
			return false, nil
		}
	}

	if linkUnit.Settings.Accessible_Ips.Active {
		found := false
		for _, val := range linkUnit.Settings.Accessible_Ips.Value {
			if val == ip {
				found = true
			}
		}
		if !found {
			return false, nil
		}
	}

	if linkUnit.Settings.Expiry_Time.Active && linkUnit.Settings.Expiry_Time.Value.Unix() < time.Now().Unix() {
		return false, nil
	}

	if linkUnit.Settings.Visits_Remaining.Active && linkUnit.Settings.Visits_Remaining.Value < 1 {
		return false, nil
	}

	return true, nil
}
