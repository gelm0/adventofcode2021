package main

import (
	"adventofcode/dive"
	"adventofcode/sonarsweep"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dayOne() {
	f, err := os.Open("sonarsweep/measurements.txt")
	if err != nil {
		fmt.Println("Failed to open input data")
	}
	defer f.Close()
	var measurements []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		measurement, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, measurement)
	}
	task1 := sonarsweep.SonarSweep(measurements)
	task2 := sonarsweep.StatSonarSweep(measurements)
	fmt.Printf("Day 1 result A: %d\n", task1)
	fmt.Printf("Day 1 result B: %d\n", task2)
}

func dayTwo() {
	f, err := os.Open("dive/input.txt")
	if err != nil {
		fmt.Println("Failed to open input data")
	}
	defer f.Close()
	var positions []dive.Position
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := strings.Split(scanner.Text(), " ")
		if len(val) != 2 {
			fmt.Println("Unable to parse dive data")
			return
		}
		direction := dive.Direction(val[0])
		distance, _ := strconv.Atoi(val[1])
		d := dive.Position{
			Direction: direction,
			Distance:  distance,
		}
		positions = append(positions, d)
	}
	task1 := dive.Dive(positions)
	task2 := dive.DiveAim(positions)

	fmt.Printf("Day 2 result A: %d\n", task1)
	fmt.Printf("Day 2 result B: %d\n", task2)
}

func main() {
	// dayOne()
	dayTwo()

}
