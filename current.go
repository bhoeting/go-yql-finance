package yql

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
)

type CurrentPiece struct {
	Ask       float64
	Open      float64
	Name      string
	Symbol    string
	Change    float64
	PrevClose float64
}

// Price returns the closing price
func (h CurrentPiece) Price() float64 {
	return h.Ask
}

// GetCurrentData returns a []CurrentPiece
// with the current data for a stock
func GetCurrentData(symbols ...string) []CurrentPiece {
	var symbolStr string
	var currentData []CurrentPiece

	for _, symbol := range symbols {
		symbolStr += "\"" + symbol + "\","
	}

	symbolStr = symbolStr[:len(symbolStr)-1]

	query := fmt.Sprintf(`SELECT * FROM %s WHERE symbol IN (%s)`,
		finaceTables["quotes"], symbolStr,
	)

	json, _ := simplejson.NewFromReader(runQuery(query))
	rows, _ := json.Get("query").Get("results").Get("quote").Array()

	for _, row := range rows {
		currentData = append(currentData, currentPieceFromRow(row))
	}

	return currentData
}

func currentPieceFromRow(row interface{}) CurrentPiece {
	c := CurrentPiece{}
	data, _ := row.(map[string]interface{})

	c.Ask, _ = strconv.ParseFloat(data["AskRealtime"].(string), 64)
	c.Open, _ = strconv.ParseFloat(data["Open"].(string), 64)
	c.Name, _ = data["Name"].(string)
	c.Symbol, _ = data["symbol"].(string)
	c.Change, _ = strconv.ParseFloat(data["Change"].(string), 64)
	c.PrevClose, _ = strconv.ParseFloat(data["PreviousClose"].(string), 64)

	return c
}
