package main
 
import (
    "fmt"
    "log"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    equity "github.com/doom160/investment-tracker/equity"
)


func main() {
    handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    var myRouter = mux.NewRouter()
    myRouter.Path("/").HandlerFunc(homePage)
    myRouter.HandleFunc("/stocks", returnStock).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func returnStock(w http.ResponseWriter, r *http.Request){
    params := r.URL.Query()
    ticker := params.Get("ticker")
    frequency := params.Get("frequency")
    strYear := params.Get("year")
    strMonth := params.Get("month")
    strDay := params.Get("day")

    var year, month, day int
    if strYear != "" {
        intYear, err := strconv.Atoi(strYear)
        if err != nil {
            fmt.Errorf("Invalid year provided %w/n", err)
        }
        if year < 0 {
            fmt.Errorf("Invalid year provided/n")
        }
    }
    if strMonth != "" {
        month, err := strconv.Atoi(strMonth)
        if err != nil {
            fmt.Errorf("Invalid month provided %w/n", err)
        }
        if month < 0 {
            fmt.Errorf("Invalid month provided/n")
        }
    }
    if strDay != "" {
        day, err := strconv.Atoi(strDay)
        if err != nil {
            fmt.Errorf("Error finding stock information %w/n", err)
        }
        if day < 0 {
            fmt.Errorf("Invalid day provided/n")
        }
    }
    
    dateRange := equity.DateRange{OffsetDay: day, OffsetMonth: month, OffsetYear: year, Frequency: frequency}

    stock, err := equity.GetEquityInfo(ticker,dateRange)
    if err != nil {
        fmt.Errorf("Error finding stock information %w", err)
    }

    b, err := json.Marshal(stock)
    if err != nil {
        fmt.Errorf("Error loading stock information %w", err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(b)
}

