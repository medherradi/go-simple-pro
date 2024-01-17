package prices

import (
	"fmt"

	"revenue.com/project/conversion"
	"revenue.com/project/filemanager"
)

type TaxIncludedPrice struct {
	TaxRate         float64
	InPrices        []float64
	TaxIncludPrices map[string]string
}

// function that read txt file
func (p *TaxIncludedPrice) LoadDataFromTxt() {
	lines, err := filemanager.ReadFilesLines("prices.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	arrayFloat, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("error occured while converting to float: " + err.Error())
		return
	}
	p.InPrices = arrayFloat
}

// function related struct
func (p *TaxIncludedPrice) Process() {
	p.LoadDataFromTxt()
	result := make(map[string]string)
	for _, val := range p.InPrices {
		calcTaxIncPrice := val * (1 + p.TaxRate)
		result[fmt.Sprintf("%.2f", val)] = fmt.Sprintf("%.2f", calcTaxIncPrice)

	}
	p.TaxIncludPrices = result
	filemanager.StoreDataAsJson(fmt.Sprintf("result_%.0f.json", p.TaxRate*100), p)
}

// adding constructor function with New key word

func NewTaxIncludPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate:  taxRate,
		InPrices: []float64{10, 20, 30},
	}
}
