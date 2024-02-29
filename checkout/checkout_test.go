package checkout

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var testPricing = map[string]Price{
    "A": {
        UnitPrice: 10,
        Special:   map[int]int{3: 25},
    },
    "B": {
        UnitPrice: 20,
        Special:   map[int]int{2: 30},
    },
}

func TestInitCheckout(t *testing.T) {
    t.Run("should initialise checkout interface with no errors", func(t *testing.T) {
        _, err := InitCheckout(testPricing)
        assert.NoError(t, err)
    })
    t.Run("should return an error when no pricing is passed", func(t *testing.T) {
        _, err := InitCheckout(nil)
        assert.Error(t, err)
    })
}
