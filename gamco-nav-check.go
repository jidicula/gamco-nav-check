package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/jidicula/go-gamco"
	"github.com/tonymackay/go-yahoo-finance"
)

func main() {
	funds, err := gamco.GetCommonFundList()
	if err != nil {
		log.Fatalf("%s", err)
	}

	navs := extractNAVs(funds)
	prices, err := extractPrices(funds)
	if err != nil {
		log.Fatalf("%s", err)
	}
	discounts, err := getDiscounts(navs, prices)
	if err != nil {
		log.Fatalf("%s", err)
	}
	today := time.Now()
	outputPath, err := dumpOutput(discounts, today)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(outputPath)
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
func getDiscount(nav string, price string) (int, error) {
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

	discountFloat, _ := discountPercent.Float64()

	discount = int(math.Round(discountFloat))
	return discount, err
}

type Stock struct {
	Symbol   string
	NAV      string
	Price    string
	Discount int
}

// getDiscounts returns a list of Stocks trading a discount relative to
// their NAV.
func getDiscounts(navs map[string]string, prices map[string]string) ([]Stock, error) {
	discountList := []Stock{}

	for k, v := range navs {
		p := prices[k]
		d, err := getDiscount(v, p)
		if err != nil {
			return discountList, err
		}

		if d >= 10 {
			s := Stock{
				Symbol:   k,
				NAV:      v,
				Price:    p,
				Discount: d,
			}
			discountList = append(discountList, s)
		}
	}

	return discountList, nil
}

// dumpOutput parses a list of Stocks and an input date, writes them into a
// temporary HTML file, and returns the path to the temp file.
func dumpOutput(sl []Stock, date time.Time) (string, error) {
	if len(sl) == 0 {
		return "", nil
	}
	dateSuffix := date.Format("2006-01-02")
	path := fmt.Sprintf("/tmp/GAMCO_%s.html", dateSuffix)

	output := fmt.Sprintf(`<html>
  <h1>Date: %s</h1>
  <table>
    <tr>
      <th>Symbol</th>
      <th>NAV (USD)</th>
      <th>Price (USD)</th>
      <th>Discount (%%)</th>
    </tr>`, dateSuffix)

	for _, s := range sl {
		row := fmt.Sprintf(`
    <tr>
      <td>%s</td>
      <td>%s</td>
      <td>%s</td>
      <td>%v</td>
    </tr>`, s.Symbol, s.NAV, s.Price, s.Discount)
		output += row
	}
	output += "\n  </table>\n</html>"

	err := os.WriteFile(path, []byte(output), 0666)
	if err != nil {
		return "", err
	}
	return path, err
}
