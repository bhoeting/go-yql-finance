package yql

import (
	"log"
	"testing"
)

// TestHistorical checks that the data
// returned from GetHistoricalData exists.
func TestHistorical(t *testing.T) {
	var date_duration [2]string = [2]string{"2016-01-01", "2016-04-01"}
	days := GetHistoricalData("GOOG", "daily", date_duration)
	weeks := GetHistoricalData("GOOG", "weekly", date_duration)
	months := GetHistoricalData("GOOG", "monthly", date_duration)

	if len(days) <= 0 || days == nil {
		log.Fatalf("%s", "Daily historical data not retrieved.\n")
	}

	if len(weeks) <= 0 || weeks == nil {
		log.Fatalf("%s", "Weekly historical data not retrieved.\n")
	}

	if len(months) <= 0 || months == nil {
		log.Fatalf("%s", "Monthly historical data not retrieved.\n")
	}
}
