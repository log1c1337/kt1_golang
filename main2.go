package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 25.4, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temperatures {

		rangeKey := int(math.Floor(temp/10) * 10)
		groups[rangeKey] = append(groups[rangeKey], temp)
	}

	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
