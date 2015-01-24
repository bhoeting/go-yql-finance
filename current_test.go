package yql

import (
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
		t.Fatal("\"days\" is nil.\n")
	}

	if len(days) == 0 {
		t.Fatal("\"days\" is empty.\n")
	}

	for index, day := range days {
		if unsafe.Sizeof(day) == 0 {
			t.Fatalf("The CurrentPiece representing %s is nil.\n", symbols[index])
		}
	}
}
