package whaletreachery

import (
	"math"
)

func findMinMax(s []int) (min, max int) {
	for _, val := range s {
		if max < val {
			max = val
		}
		if min > val {
			min = val
		}
	}
	return
}

func countMeanSides(s []int, mean int) (minCount, maxCount int) {
	for _, val := range s {
		if mean < val {
			maxCount += 1
		} else {
			minCount += 1
		}
	}
	return
}

func calculateExtraFuel(steps int) (extraFuel int) {
	i := 1
	for i < steps {
		extraFuel += i
		i += 1
	}
	return
}

// Horrible complexity, divide and conquer algorithm with O(logn) would probably
// be feasible
func searchMinDistanceLeft(crabPositions []int, mean int, extraFuelStep bool) (minFuel int) {
	var previousDistance = math.MaxInt64
	for mean >= 0 {
		for _, val := range crabPositions {
			steps := int(math.Abs(float64(val - mean)))
			minFuel += steps
			if extraFuelStep {
				minFuel += calculateExtraFuel(steps)
			}
		}
		if minFuel < previousDistance {
			previousDistance = minFuel
		}
		mean -= 1
		minFuel = 0
	}
	return previousDistance
}

func searchMinDistanceRight(crabPositions []int, mean, max int, extraFuelStep bool) (minFuel int) {
	var previousDistance = math.MaxInt64
	for mean <= max {
		for _, val := range crabPositions {
			steps := int(math.Abs(float64(val - mean)))
			minFuel += steps
			if extraFuelStep {
				minFuel += calculateExtraFuel(steps)
			}
		}
		if minFuel < previousDistance {
			previousDistance = minFuel
		}
		mean += 1
		minFuel = 0
	}
	return previousDistance
}

// This is basically just a guessing algorithm with pretty bad complexity
// It works to solve this problem but I can't gurantee it's correctness
// It works by taking the average number and then looking which side of that number
// contains the most amount of numbers.
// It then performs a search over those numbers basically making it O(N/2) instead of O(N)
// Probably some O(logn) solution prob divide and conq would make better fit here
// Note: You can guarantee correctness by commenting out the if else case in the end
// and compare both values to each other

func startCrabDance(crabPositions []int, extraFuelStep bool) int {
	min, max := findMinMax(crabPositions)
	meanMinMax := int(math.Abs(float64((max - min) / 2)))
	minCount, maxCount := countMeanSides(crabPositions, meanMinMax)
	guess := 0
	if minCount > maxCount {
		guess = searchMinDistanceLeft(crabPositions, meanMinMax, extraFuelStep)
	} else {
		guess = searchMinDistanceRight(crabPositions, meanMinMax, max, extraFuelStep)
	}
	return guess
}

func CrabDance(crabPositions []int) int {
	return startCrabDance(crabPositions, false)
}

func AdvancedCrabDancing(crabPositions []int) int {
	return startCrabDance(crabPositions, true)
}
