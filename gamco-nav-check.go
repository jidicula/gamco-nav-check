package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/jidicula/go-gamco"
	"github.com/tonymackay/go-yahoo-finance"
)

func main() {
	fmt.Println("Hello, world!")
}

// getNAVs returns a map of NAVs for each GAMCO common stock.
func extractNAVs(fl []gamco.Fund) map[string]string {
	navs := make(map[string]string)
	for _, v := range fl {
		navs[v.Symbol] = v.NAV
	}
	return navs
}

// extractPrices returns a map of market price strings for each GAMCO common
// stock by querying the Yahoo Finance API.
func extractPrices(fl []gamco.Fund) (map[string]string, error) {
	prices := make(map[string]string)
	for _, v := range fl {
		symbol := v.Symbol
		result, err := yahoo.Quote(symbol)
		if err != nil {
			return prices, err
		}
		prices[symbol] = result.QuoteSummary.Result[0].Price.RegularMarketPrice.Fmt
	}

	return prices, nil
}

// getDiscount returns the rounded price discount: (NAV/price - 1) * 100.
func getDiscount(price string, nav string) (int, error) {
	var err error
	var discount int
	p := new(big.Rat)
	p.SetString(price)

	n := new(big.Rat)
	n.SetString(nav)

	// Handle any panic that occurs from zero division attempt
	defer func() {
		if recover() != nil {
			err = errors.New("Price is 0, cannot calculate discount")
		}
	}()

	q := new(big.Rat)
	q.Quo(n, p)

	discountProportion := new(big.Rat)
	discountProportion.Sub(q, big.NewRat(1, 1))

	discountPercent := new(big.Rat)
	discountPercent.Mul(discountProportion, big.NewRat(100, 1))

	discountFloat, ok := discountPercent.Float64()
	if !ok {
		return discount, fmt.Errorf("Could not convert discount to float: %v", discountFloat)
	}

	discount = int(math.Round(discountFloat))
	return discount, err
}
