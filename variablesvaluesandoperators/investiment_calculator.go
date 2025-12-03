package main

import (
	"fmt"
	"math"
)

func main() {
	var investimentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = investimentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Printf("Future value: %.2f\n", futureValue)
}
