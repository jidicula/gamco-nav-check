package main

import (
	"fmt"

	"github.com/jidicula/go-gamco"
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
