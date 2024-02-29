package main

import (
    "fmt"
    "github.com/Caitlin-B/checkout-kata/checkout"
)

var scannedItems = []string{"B", "A", "B"}

func main() {
    // read prices csv
    prices, err := checkout.ScanPrices("prices.csv")
    if err != nil {
        fmt.Println(err)
        return
    }

    // init checkout with prices set
    co, err := checkout.InitCheckout(prices)
    if err != nil {
        fmt.Println(err)
        return
    }
    // scan items
    for _, item := range scannedItems {
        co.Scan(item)
    }
    // get total
    fmt.Println(co.GetTotalPricing())
}
