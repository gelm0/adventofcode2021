package sonarsweep

import "fmt"

func sum(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func SonarSweep(measurements []int) int {
	counter := 0
	for i, _ := range measurements {
		if i == len(measurements)-1 {
			break
		}
		if measurements[i+1]-measurements[i] > 0 {
			counter += 1
		}

	}
	return counter
}

func StatSonarSweep(measurements []int) int {
	counter := 0
	for i, _ := range measurements {
		if i == len(measurements)-3 {
			break
		}
		window1 := sum(measurements[i : i+3])
		window2 := sum(measurements[i+1 : i+4])
		fmt.Println(window1)
		fmt.Println(window2)
		fmt.Println(counter)
		fmt.Println("---")
		if window2-window1 > 0 {
			counter += 1
		}
	}
	return counter
}
