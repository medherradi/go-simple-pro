package prices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPrice struct {
	TaxRate         float64
	InPrices        []float64
	TaxIncludPrices map[string][]float64
}

// function that read txt file
func (p TaxIncludedPrice) LoadDataFromTxt() ([]float64, error) {
	file, err := os.Open("prices.txt")
	if err != nil {
		return nil, errors.New("something went wrong: " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	var arrayFloat []float64
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return nil, errors.New("reading content in file failed: " + err.Error())
	}
	for _, value := range lines {
		floatNum, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, errors.New("can not parse to float: " + err.Error())
		}
		arrayFloat = append(arrayFloat, floatNum)
	}
	return arrayFloat, nil
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
