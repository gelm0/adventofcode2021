package extendedpoly

import (
	"math"
)

func Step(steps int, polyChain map[string]int, polyConv map[string]string) map[string]int {
	newChain := make(map[string]int)
	for k, v := range polyChain {
		conversionValue := polyConv[k]
		pair1 := string(k[0]) + conversionValue
		pair2 := conversionValue + string(k[1])
		newChain[pair1] += v
		newChain[pair2] += v
	}
	steps -= 1
	if steps == 0 {
		return newChain
	}
	return Step(steps, newChain, polyConv)
}

func CountOcurrence(polyChain map[string]int, lastValue string) int {
	occurence := make(map[string]int)
	var maxValue int
	minValue := math.MaxInt
	for k, v := range polyChain {
		char1 := k[0]
		occurence[string(char1)] += v
	}
	// The polymerchain starts and ends with the same character. We don't keep track on order that so we
	// need to add it after
	occurence[lastValue] += 1
	for _, v := range occurence {
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
	}
	return maxValue - minValue
}
