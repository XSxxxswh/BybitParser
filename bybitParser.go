package bybitParser

import (
	"fmt"
	"math"
)

func ExchangeRate() {
	var awerage float64
	prices, err := GetRate("RUB", "100000", "10")
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

	fmt.Println(math.Round(awerage*100) / 100)
}
