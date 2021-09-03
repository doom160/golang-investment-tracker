package main
 
import (
    "flag"
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "github.com/fxtlabs/date"
	"github.com/leekchan/accounting"
    "github.com/piquette/finance-go/quote"
    "github.com/sirupsen/logrus"
    "github.com/piquette/finance-go/chart"
    "github.com/piquette/finance-go/datetime"
)
// https://query2.finance.yahoo.com/v8/finance/chart/AAPL?symbol=AAPL&period1=1605796200&period2=99999999999999&interval=1d
/*

func getEpochTime() int64 {
    return time.Today().AddDate(0,-6,0)Unix()
}


*/


func main() {
    flag.Parse()
 
    if len(flag.Args()) == 0 {
        logrus.Fatalf("No argument provided, exected at least one stock symbol. Example: %v cldr goog aapl intc amd ...", os.Args[0])
    }
 
    cf := accounting.Accounting{Symbol: "$", Precision: 2}
    smbls := flag.Args()
 
    iter := quote.List(smbls)
    epoch := date.Today().UTC().AddDate(0,-6,0).Unix()

    resp, err := http.Get(fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/AAPL?symbol=AAPL&period1=%d&period2=99999999999999&interval=1d", epoch))

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    sb := string(body)
    fmt.Printf(sb)
    for iter.Next() {
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
	 }
}