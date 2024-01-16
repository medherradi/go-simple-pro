package prices

type TaxIncludedPrice struct {
	TaxRate         float64
	InPrices        []float64
	TaxIncludPrices map[string][]float64
}
