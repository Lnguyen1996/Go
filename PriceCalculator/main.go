package main

import (
	"example.com/price_calculator/cmdmanager"
	"example.com/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// outputPath := fmt.Sprintf("result_%.0f.json", taxRate*100)

		// fm := filemanager.New("prices.txt",outputPath)

		cmdm := cmdmanager.New()

		pricejob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)

		pricejob.Process()
	}

}
