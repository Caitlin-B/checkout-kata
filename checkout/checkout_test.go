package checkout

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var testPricing = map[string]*Price{
    "A": {
        UnitPrice:    10,
        SpecialPrice: 25,
        SpecialCount: 3,
    },
    "B": {
        UnitPrice:    20,
        SpecialPrice: 30,
        SpecialCount: 2,
    },
    "C": {
        UnitPrice: 25,
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

func TestCheckout_Scan(t *testing.T) {
    t.Run("should scan inputs and add them to checkout object items list", func(t *testing.T) {
        co := Checkout{}
        co.Scan("A")
        co.Scan("B")
        co.Scan("A")
        co.Scan("C")
        assert.ElementsMatch(t, []string{"A", "B", "A", "C"}, co.Items)
    })
}

func TestCheckout_GetTotalPricing(t *testing.T) {
    cases := []struct {
        items    []string
        expected int
    }{
        {
            items:    []string{"A"},
            expected: 10,
        },
        {
            items:    []string{"A", "A", "A", "B"},
            expected: 45,
        },
        {
            items:    []string{"A", "A", "A", "A", "B"},
            expected: 55,
        },
        {
            items:    []string{"A", "B", "C", "A", "B"},
            expected: 75,
        },
    }

    for _, tc := range cases {
        t.Run("should calculate total pricing correctly", func(t *testing.T) {
            co := Checkout{
                Items:  tc.items,
                Prices: testPricing,
            }
            assert.Equal(t, tc.expected, co.GetTotalPricing())
        })
    }

}
