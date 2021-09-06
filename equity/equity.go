package equity

import (
    "flag"
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/fxtlabs/date"
//	"github.com/leekchan/accounting"
)
// https://query2.finance.yahoo.com/v8/finance/chart/AAPL?symbol=AAPL&period1=1605796200&period2=99999999999999&interval=1d
/*

func getEpochTime() int64 {
    return time.Today().AddDate(0,-6,0)Unix()
}

https://www.reddit.com/r/sheets/wiki/apis/finance#wiki_finance_apis
https://query2.finance.yahoo.com/v10/finance/quoteSummary/NVDA?modules=defaultKeyStatistics%2CassetProfile%2CtopHoldings%2CfundPerformance%2CfundProfile%2CesgScores&ssl=true


*/


func getHistoricalData(string ticker) {

    epoch := date.Today().UTC().AddDate(0,-6,0).Unix()

    resp, err := http.Get(fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=99999999999999&interval=1d",ticker, epoch))

    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    sb := string(body)
    //fmt.Printf(sb)

    myStoredVariable := HistoryData{}
    json.Unmarshal([]byte(sb), &myStoredVariable)
    fmt.Printf("%f", myStoredVariable.HistoryData.Result[0].Meta.RegularMarketPrice)
}



type HistoryData struct {
	HistoryData Chart `json:"chart"`
}
 
type Chart struct {
	Result []Result `json: "result"`
	Error  string   `json: "error"`
}

type Result struct {
	Meta Meta `json: "meta"`
	Timestamp  []int32    `json: "timestamp"`
    Indictators Indictators `json: "indicators"`
}

type Meta struct {
	Currency string `json: "currency"`
	Symbol  string `json: "symbol"`
    RegularMarketPrice float32 `json: "regularMarketPrice"`
}

type Indictators struct {
    Quote Quote `json: "quote"`
}

type Quote struct {
    Open []float32 `json: "open"`
    High []float32 `json: "high"`
    Low []float32 `json: "low"`
    Close []float32 `json: "close"`
    volume []int32 `json: "volume"`
}

