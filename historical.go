package yql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
)

// HistoricalPiece is a struct
// representation of the data
// for a trading day returned
// from the YQL API
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
// data_duration[2] = []string{"2016-01-01", "2016-02-01"}
func GetHistoricalData(symbol string, timeInterval string, date_duration [2]string) []HistoricalPiece {
	date_begin, _ := time.Parse("2006-01-02", date_duration[0])
	date_end, _ := time.Parse("2006-01-02", date_duration[1])
	date_begin_str := fmt.Sprintf("&a=%02d&b=%02d&c=%d", date_begin.Month()-1, date_begin.Day(), date_begin.Year())
	date_end_str := fmt.Sprintf("&d=%02d&e=%02d&f=%d", date_end.Month()-1, date_end.Day(), date_end.Year())

	var historicalData []HistoricalPiece

	switch timeInterval {
	case "daily":
		timeInterval = "d"
	case "weekly":
		timeInterval = "w"
	case "monthly":
		timeInterval = "m"
	}

	/*
		query := fmt.Sprintf(
			`SELECT * FROM %s WHERE url='%s%s' AND
			columns='Date,Open,High,Low,Close,Volume,AdjClose'`,
			finaceTables["historical"], historicalURL+symbol, "&g="+timeInterval)
	*/
	query2 := fmt.Sprintf(
		`SELECT * FROM %s WHERE url='%s%s%s%s' AND
		columns='Date,Open,High,Low,Close,Volume,AdjClose'`,
		finaceTables["historical"], historicalURL+symbol, "&g="+timeInterval, date_begin_str, date_end_str)

	json, _ := simplejson.NewFromReader(runQuery(query2))

	rows, _ := json.Get("query").Get("results").Get("row").Array()

	for index, row := range rows {
		// The first row must be skipped since it contains table headers
		if index == 0 {
			continue
		}

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
	h.Date, _ = time.Parse("2006-01-02", data["Date"].(string))
	h.Volume, _ = strconv.ParseFloat(data["Volume"].(string), 64)
	h.AdjClose, _ = strconv.ParseFloat(data["AdjClose"].(string), 64)

	return h
}
