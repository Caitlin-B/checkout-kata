package checkout

import (
    "encoding/csv"
    "os"
    "strconv"
    "strings"
)

// ScanPrices scans prices from a csv formats them into a map of SKU > Price
func ScanPrices(path string) (map[string]*Price, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    r := csv.NewReader(file)
    p, err := r.ReadAll()
    if err != nil {
        return nil, err
    }

    prices := make(map[string]*Price, len(p))

    for _, item := range p {
        unitP, err2 := strconv.Atoi(item[1])
        if err2 != nil {
            return nil, err2
        }
        prices[item[0]] = &Price{
            UnitPrice: unitP,
        }
        if item[2] == "" {
            continue
        }
        special := strings.Split(item[2], " ")
        specialC, err2 := strconv.Atoi(special[0])
        if err2 != nil {
            return nil, err2
        }
        specialP, err2 := strconv.Atoi(special[2])
        if err2 != nil {
            return nil, err2
        }
        prices[item[0]].SpecialPrice = specialP
        prices[item[0]].SpecialCount = specialC
    }

    return prices, nil
}
