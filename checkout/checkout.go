package checkout

import (
    "errors"
)

// ICheckout interface to represent
type ICheckout interface {
    Scan(input string)
    GetTotalPricing() int
}

// Price represents information on an items pricing
type Price struct {
    // UnitPrice the standard price
    UnitPrice int
    // SpecialPrice the price when an item is on offer
    SpecialPrice int
    // SpecialCount the number of items required to trigger the special price
    SpecialCount int
}

// Checkout meets the ICheckout interface and stores information what items have been scanned and the list of prices
type Checkout struct {
    // Items list of items that have been scanned
    Items []string
    // Prices map of SKU > Price for all available items
    Prices map[string]*Price
}

// Scan adds and item to the list of items on a checkout
func (c *Checkout) Scan(item string) {
    c.Items = append(c.Items, item)
}

// GetTotalPricing returns the total price of all scanned items
func (c *Checkout) GetTotalPricing() int {
    count := map[string]int{}

    for _, item := range c.Items {
        count[item]++
    }

    tot := 0
    for sku, i := range count {
        price := c.Prices[sku]
        // rem is number of items not covered by a special offer
        rem := i
        // calc exceptions
        if price.SpecialCount != 0 {
            rem = i % price.SpecialCount
            tot += (i - rem) / price.SpecialCount * price.SpecialPrice
        }
        // calc remaining
        tot += rem * price.UnitPrice

    }

    return tot
}

func InitCheckout(prices map[string]*Price) (ICheckout, error) {
    if prices == nil {
        return nil, errors.New("prices are required to initialise checkout")
    }

    return &Checkout{
        Items:  nil,
        Prices: prices,
    }, nil
}
