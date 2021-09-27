package equity

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
    "github.com/fxtlabs/date"
)

func verifyDateRange(frequency string) bool {
    switch frequency {
        case 
            "1d", 
            "5d", 
            "1mo", 
            "3mo", 
            "6mo", 
            "1y", 
            "2y", 
            "5y", 
            "10y", 
            "ytd",
            "max":
        return true
    }
    return false 
}

func GetEquityInfo(ticker string, dateRange DateRange) (equityInfo EquityInfo, err error) {

    // Validate ticker
    ticker = strings.ToUpper(strings.TrimSpace(ticker))
    if ticker == "" {
        return equityInfo, errors.New("ticker should not be empty")
    }

    // Validate dateRange
    if dateRange.OffsetDay < 0 || dateRange.OffsetMonth < 0 || dateRange.OffsetYear < 0 {
        return equityInfo, errors.New("DateRange OffsetDay, OffsetMonth, OffsetYear should be positive value")
    }
    if !verifyDateRange(dateRange.Frequency){
        return equityInfo, errors.New("dateRange.Frequency are not valid. {1d, 5d, 1mo, 3mo, 6mo, 1y, 2y, 5y, 10y, ytd, max}")
    }

    var equity Equity
    epoch := date.Today().UTC().AddDate(-dateRange.OffsetDay, -dateRange.OffsetMonth, -dateRange.OffsetYear).Unix()
    currentDate := date.Today().UTC().Unix()

    var resp *http.Response
    resp, err = http.Get(fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=%d&interval=%s",ticker, epoch, dateRange.Frequency, currentDate))

    if err != nil {
        return equityInfo, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return equityInfo, err
    }

    json.Unmarshal([]byte(string(body)), &equity)
    //fmt.Printf("%f", myStoredVariable.Chart.Result[0].Meta.RegularMarketPrice)
    equityInfo = EquityInfo{ Symbol:                equity.Chart.Result[0].Meta.Symbol, 
                             Currency:              equity.Chart.Result[0].Meta.Currency,
                             RegularMarketPrice:    equity.Chart.Result[0].Meta.RegularMarketPrice,
                             ChartPreviousClose:    equity.Chart.Result[0].Meta.ChartPreviousClose,
                             Timestamp:             equity.Chart.Result[0].Timestamp,
                             Open:                  equity.Chart.Result[0].Indicators.Quote[0].Open,
                             High:                  equity.Chart.Result[0].Indicators.Quote[0].High,
                             Low:                   equity.Chart.Result[0].Indicators.Quote[0].Low,
                             Close:                 equity.Chart.Result[0].Indicators.Quote[0].Close,
                             Volume:                equity.Chart.Result[0].Indicators.Quote[0].Volume}
                             
    return equityInfo, nil
}

type DateRange struct {
    OffsetDay int32
    OffsetMonth int32
    OffsetYear int32
    Frequency string
}

type Equity struct {
	Chart Chart `json:"chart"`
}
 
type Chart struct {
	Result []Result `json: "result"`
	Error  string   `json: "error"`
}

type Result struct {
	Meta Meta `json: "meta"`
	Timestamp  []int32    `json: "timestamp"`
    Indicators Indicators `json: "indicators"`
}

type Meta struct {
	Currency string `json: "currency"`
	Symbol  string `json: "symbol"`
    RegularMarketPrice float32 `json: "regularMarketPrice"`
    ChartPreviousClose float32 `json: "chartPreviousClose"`
}

type Indicators struct {
    Quote []Quote `json: "quote"`
}

type Quote struct {
    Open []float32 `json: "open"`
    High []float32 `json: "high"`
    Low []float32 `json: "low"`
    Close []float32 `json: "close"`
    Volume []int32 `json: "volume"`
}

type EquityInfo struct {
    Symbol  string `json: "symbol"`
    Currency string `json: "currency"`
    RegularMarketPrice float32 `json: "regularMarketPrice"`
    ChartPreviousClose float32 `json: "chartPreviousClose"`
    Timestamp  []int32    `json: "timestamp"`
    Open []float32 `json: "open"`
    High []float32 `json: "high"`
    Low []float32 `json: "low"`
    Close []float32 `json: "close"`
    Volume []int32 `json: "volume"`
}


