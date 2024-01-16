package main

import (
	"revenue.com/project/prices"
)

func main() {

	taxRates := []float64{0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludPrice(taxRate)
		priceJob.Process()
	}

}
