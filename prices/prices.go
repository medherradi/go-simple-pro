package prices

import (
	"bufio"
	"fmt"
	"os"

	"revenue.com/project/conversion"
)

type TaxIncludedPrice struct {
	TaxRate         float64
	InPrices        []float64
	TaxIncludPrices map[string][]float64
}

// function that read txt file
func (p *TaxIncludedPrice) LoadDataFromTxt() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("error occured while opening file: " + err.Error())
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading content in file failed: " + err.Error())
		file.Close()
		return
	}
	arrayFloat, err := conversion.StringsToFloat(lines)
	if err != nil {
		file.Close()
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
	fmt.Println(result)
}

// adding constructor function with New key word

func NewTaxIncludPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate:  taxRate,
		InPrices: []float64{10, 20, 30},
	}
}
