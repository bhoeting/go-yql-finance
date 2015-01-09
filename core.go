package yql

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type TradingDay interface {
	Price() float64
}

const (
	baseUrl       string = "query.yahooapis.com"
	publicApiUrl  string = "http://query.yahooapis.com/v1/public/yql"
	datatablesUrl string = "store://datatables.org/alltableswithkeys"
	historicalUrl string = "http://ichart.finance.yahoo.com/table.csv?s="
)

var finaceTables map[string]string = map[string]string{
	"historical": "csv",
	"quotes":     "yahoo.finance.quotes",
}

// runQuery runs the query and retuns the
// results in an io.Reader
func runQuery(query string) io.Reader {
	queryUrl := buildUrl(query)
	return getJSON(queryUrl)
}

// buildUrl creates a YQL URL from a query
func buildUrl(query string) string {
	params := url.Values{}
	params.Add("q", query)
	params.Add("format", "json")
	params.Add("env", datatablesUrl)

	return publicApiUrl + "?" + params.Encode()
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
