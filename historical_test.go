package yql

import (
	"log"
	"testing"
)

// TestHistorical checks that the data
// returned from GetHistoricalData exists.
func TestHistorical(t *testing.T) {
	days := GetHistoricalData("GOOG", "daily")
	weeks := GetHistoricalData("GOOG", "weekly")
	months := GetHistoricalData("GOOG", "monthly")

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
