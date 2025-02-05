package main

import "fmt"

func main() {

	const eurToUsd float64 = 1.04
	const usdToRub float64 = 99.20

	// курсы на 5 фев 2025

	const eur float64 = 500
	var result = eur * eurToUsd * usdToRub
	fmt.Print(result)

}
