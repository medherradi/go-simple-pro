package prices

import (
	"fmt"
	"strconv"
)

type TaxIncludedPrice struct {
	TaxRate         float64
	InPrices        []float64
	TaxIncludPrices map[string][]float64
}

// function related struct
func (p TaxIncludedPrice) Process() {
	result := make(map[string]float64)

	for _, val := range p.InPrices {

		result[strconv.Itoa(int(val))] = val * (1 + p.TaxRate)

	}
	fmt.Println(result)
}

// adding constructor function with New key word

func NewTaxIncludPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate:  taxRate,
		InPrices: []float64{10, 20, 30},
	}
}
