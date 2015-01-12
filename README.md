# go-yql-finance
Yahoo finance API wrapper for Go

## Usage
```go
import (
  "github.com/bhoeting/go-yql-finance"
)
```

### Historical Data
`GetHistoricalData()` returns a slice of type `HistoricalPiece` which is defined below:
```go
type HistoricalPiece struct {
	Low      float64
	High     float64
	Open     float64
	Close    float64
	Volume   float64
	AdjClose float64
	Date     time.Time
}
```
The function can be used like this:
```go
// The second parameter is the time interval.
// You can pass in "d", "w", "m", "daily", "weekly", or "montly".
days := yql.GetHistoricalData("GOOG", "daily")

for _, day := range days {
	fmt.Println(day.Close)
}
```

### Current Data
`GetCurrentData()` returns a slice of type `CurrentPiece` which is defined below:
```go
type CurrentPiece struct {
	Ask       float64
	Open      float64
	Name      string
	Symbol    string
	Change    float64
	PrevClose float64
}
```
The function can be used like this:
```go
current := yql.GetCurrentData("GOOG", "AAPL")

for _, day := range days {
	fmt.Println(day.Ask)
}
```
### TradingDay
Both the `CurrentPiece` and `HistoricalPeice` types implement the interface `TradingDay`
```go
type TradingDay interface {
	Price() float64
}
```
`Price()` returns the `Ask` field from the `CurrentPiece` and the `Close` field from the `HistoricalPiece`

