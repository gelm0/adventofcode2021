package main

import (
	"adventofcode/binarydiagnostic"
	"adventofcode/dive"
	"adventofcode/sonarsweep"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData(inputFile string) *os.File {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input data")
	}
	return f
}

func dayOne() {
	var measurements []int
	file := getInputData("sonarsweep/measurements.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
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
	var positions []dive.Position
	file := getInputData("dive/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
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

func dayThree() {
	file := getInputData("binarydiagnostic/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var binaryData []string
	for scanner.Scan() {
		binaryData = append(binaryData, scanner.Text())
	}
	task1 := binarydiagnostic.Diagnostic(binaryData)
	co2 := binarydiagnostic.DiagnoseOxygen(binaryData)
	oxygen := binarydiagnostic.DiagnoseCO2(binaryData)
	task2 := oxygen * co2
	fmt.Printf("Day 3 result A: %d\n", task1)
	fmt.Printf("Day 3 result B: %d\n", task2)

}

func main() {
	//dayOne()
	//dayTwo()
	dayThree()
}
