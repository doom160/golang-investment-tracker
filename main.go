package main
 
import (
    "fmt"
    equity "github.com/doom160/investment-tracker/equity"
)


func main() {
    stock := equity.GetEquity("AAPL") 
    fmt.Println(stock.Chart.Result[0].Meta.RegularMarketPrice)
}



