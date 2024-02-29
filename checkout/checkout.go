package checkout

import (
    "errors"
)

type ICheckout interface {
    Scan(input string)
    GetTotalPricing() int
}

type Price struct {
    UnitPrice    int
    SpecialPrice int
    SpecialCount int
}

type Checkout struct {
    Items  []string
    Prices map[string]*Price
}

func (c *Checkout) Scan(item string) {
    c.Items = append(c.Items, item)
}

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
