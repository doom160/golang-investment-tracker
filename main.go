package main
 
import (
    "fmt"
    equity "github.com/doom160/investment-tracker/equity"
)


func main() {
    stock, err := equity.GetEquity("AAPL")
    if err != nil {
        fmt.Errorf("Error loading stock information %w", err)
    }
    fmt.Println(stock.Chart.Result[0].Meta.RegularMarketPrice)
}


