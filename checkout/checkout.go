package checkout

import "errors"

type ICheckout interface {
    Scan(input string)
    GetTotalPricing() int
}

type Price struct {
    UnitPrice int
    Special   map[int]int
}

type Checkout struct {
    Items  []string
    Prices map[string]Price
}

func (c Checkout) Scan(item string) {}

func (c Checkout) GetTotalPricing() int {
    return 0
}

func InitCheckout(prices map[string]Price) (ICheckout, error) {
    if prices == nil {
        return nil, errors.New("prices are required to initialise checkout")
    }

    return Checkout{
        Items:  nil,
        Prices: prices,
    }, nil
}
