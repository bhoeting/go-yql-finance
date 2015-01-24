package yql

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// TradingDay has a method
// Price() that returns the
// most useful money value
// from its implementation
type TradingDay interface {
	Price() float64
}

const (
	baseURL       string = "query.yahooapis.com"
	publicAPIURL  string = "http://query.yahooapis.com/v1/public/yql"
	datatablesURL string = "store://datatables.org/alltableswithkeys"
	historicalURL string = "http://ichart.finance.yahoo.com/table.csv?s="
)

var finaceTables = map[string]string{
	"historical": "csv",
	"quotes":     "yahoo.finance.quotes",
}

// runQuery runs the query and retuns the
// results in an io.Reader
func runQuery(query string) io.Reader {
	queryURL := buildURL(query)
	return getJSON(queryURL)
}

// buildURL creates a YQL URL from a query
func buildURL(query string) string {
	params := url.Values{}
	params.Add("q", query)
	params.Add("format", "json")
	params.Add("env", datatablesURL)

	return publicAPIURL + "?" + params.Encode()
}

// getJSON returns the JSON response
// from a request to a given URL
func getJSON(url string) io.Reader {
	resp, errReq := http.Get(url)

	if errReq != nil {
		log.Fatalf("%s", errReq.Error())
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return bytes.NewReader(body)
}
