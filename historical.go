package yql

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
	"time"
)

type HistoricalPiece struct {
	Low      float64
	High     float64
	Open     float64
	Close    float64
	Volume   float64
	AdjClose float64
	Date     time.Time
}

// Price returns the closing price
func (h HistoricalPiece) Price() float64 {
	return h.Close
}

// GetHistoricalData returns a []HistoricalPiece
// of a stock's historical data
func GetHistoricalData(symbol string) []HistoricalPiece {
	var historicalData []HistoricalPiece

	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE url='%s%s' AND 
		columns='Date,Open,High,Low,Close,Volume,AdjClose'`,
		finaceTables["historical"], historicalUrl+symbol, "&g=m")

	json, _ := simplejson.NewFromReader(runQuery(query))
	rows, _ := json.Get("query").Get("results").Get("row").Array()

	for _, row := range rows {
		historicalData = append(historicalData, newHistoricalPieceFromRow(row))
	}

	return historicalData
}

// newHistoricalPieceFromRow returns a
func newHistoricalPieceFromRow(row interface{}) HistoricalPiece {
	h := HistoricalPiece{}
	data, _ := row.(map[string]interface{})

	h.Low, _ = strconv.ParseFloat(data["Low"].(string), 64)
	h.High, _ = strconv.ParseFloat(data["High"].(string), 64)
	h.Open, _ = strconv.ParseFloat(data["Open"].(string), 64)
	h.Close, _ = strconv.ParseFloat(data["Close"].(string), 64)
	h.Date, _ = time.Parse("2015-01-02", data["Date"].(string))
	h.Volume, _ = strconv.ParseFloat(data["Volume"].(string), 64)
	h.AdjClose, _ = strconv.ParseFloat(data["AdjClose"].(string), 64)

	return h
}
