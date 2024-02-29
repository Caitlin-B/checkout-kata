package checkout

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var testPrices = map[string]*Price{
    "A": {
        UnitPrice:    50,
        SpecialPrice: 130,
        SpecialCount: 3,
    },
    "B": {
        UnitPrice: 20,
    },
}

func TestScanPrices(t *testing.T) {
    t.Run("should scan and format price list from csv", func(t *testing.T) {
        got, err := ScanPrices("./testData/prices_test.csv")
        assert.NoError(t, err)
        assert.Equal(t, testPrices["A"], got["A"])
        assert.Equal(t, testPrices["B"], got["B"])

    })
}
