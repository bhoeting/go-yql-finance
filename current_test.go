package yql

import (
	"log"
	"testing"
	"unsafe"
)

// TestCurrent tests that the data returned from
// GetCurrentData() is not nil, empty, and
// each HistoricalPiece is not nil.
func TestCurrent(t *testing.T) {
	var symbols []string = []string{"GOOG", "AAPL", "SPY"}
	days := GetCurrentData(symbols...)

	if days == nil {
		log.Fatal("\"days\" is nil.\n")
	}

	if len(days) == 0 {
		log.Fatal("\"days\" is empty.\n")
	}

	for index, day := range days {
		if unsafe.Sizeof(day) == 0 {
			log.Fatalf("The CurrentPiece representing %s is nil.\n", symbols[index])
		}
	}
}
