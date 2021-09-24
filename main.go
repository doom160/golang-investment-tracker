package main
 
import (
    "fmt"
    "log"
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
    myRouter.Path("/stocks").Queries("ticker","{ticker}").HandlerFunc(returnStock)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func returnStock(w http.ResponseWriter, r *http.Request){
    ticker := r.FormValue("ticker")
    stock, err := equity.GetEquityInfo(ticker)
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

