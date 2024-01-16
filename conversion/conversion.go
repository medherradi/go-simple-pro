package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(arrayString []string) ([]float64, error) {
	arrayFloat := make([]float64, len(arrayString))
	for index, value := range arrayString {
		floatNum, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("can not parse to float: " + err.Error())
			return []float64{}, errors.New("can not parse to float: " + err.Error())
		}
		arrayFloat[index] = floatNum
	}
	return arrayFloat, nil
}
