package lanternfish

func sumSlice(s []int) (sum int) {
	for _, val := range s {
		sum += val
	}
	return
}

// Newborn babies are set to 8
// Parents are reset to 6
func calculateFishReproduction(fishPopulation []int, days int) int {
	dailyPopulation := fishPopulation[1:9]
	dailyPopulation = append(dailyPopulation, fishPopulation[0])
	dailyPopulation[6] += fishPopulation[0]
	days -= 1
	if days == 0 {
		return sumSlice(dailyPopulation)
	}
	return calculateFishReproduction(dailyPopulation, days)
}

func GetNumberOfFish(fishPopulation []int, days int) int {
	return calculateFishReproduction(fishPopulation, days)
}
