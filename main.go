package main

import "github.com/Caitlin-B/checkout-kata/checkout"

var scannedItems = []string{"B", "A", "B"}

func main() {
    // read prices csv
    // init checkout with prices set
    // scan items
    // get total
    t := make(map[string]checkout.Price, 0)
    checkout.InitCheckout(t)
}
