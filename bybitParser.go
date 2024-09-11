package bybitParser

import (
	"fmt"
	"math"
)

func ExchangeRate(curr, amount, size) float64{
	var awerage float64
	prices, err := GetRate(curr, amount, size)
	if err != nil {
		panic(err)
	}
	var sum float64
	if ((prices[0]*100 - prices[4]) - 100) < 1 {

		for i := 0; i < 5; i++ {
			sum += prices[i]
		}

	} else {

		for i := 5; i < 10; i++ {
			sum += prices[i]
		}
	}
	awerage = sum / 5

	return awerage
}
