package main

import (
	"fmt"

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
