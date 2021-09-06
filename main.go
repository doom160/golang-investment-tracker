package main
 
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


func main() {
    flag.Parse()
 
    if len(flag.Args()) == 0 {
        fmt.Printf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
    }
 
    //cf := accounting.Accounting{Symbol: "$", Precision: 2}
    //smbls := flag.Args()
 
    //iter := quote.List(smbls)
    epoch := date.Today().UTC().AddDate(0,-6,0).Unix()

    resp, err := http.Get(fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/AAPL?symbol=AAPL&period1=%d&period2=99999999999999&interval=1d", epoch))

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

    /*for iter.Next() {
        q := iter.Quote()
        fmt.Printf("------- %v -------\n", q.ShortName)
        fmt.Printf("Current Price: %v\n", cf.FormatMoney(q.Ask))
        fmt.Printf("52wk High: %v\n", cf.FormatMoney(q.FiftyTwoWeekHigh))
        fmt.Printf("52wk Low: %v\n", cf.FormatMoney(q.FiftyTwoWeekLow))
        fmt.Printf("-----------------\n")
    }

	 f, err := quote.Get("SPY")
	 if err != nil {
	 	fmt.Println(err)
	 } else {
	 	fmt.Println(f)
	 }

	 params := &chart.Params{
	 	Symbol:   "TWTR",
	 	Interval: datetime.OneHour,
	 }
	 iter2 := chart.Get(params)
	
	 for iter2.Next() {
	 	b := iter2.Bar()
	 	fmt.Println(datetime.FromUnix(b.Timestamp))
	
	 }
	 if iter2.Err() != nil {
	 	fmt.Println(iter2.Err())
	 }*/
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

