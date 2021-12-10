package main

import (
	"adventofcode/binarydiagnostic"
	"adventofcode/bingo"
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

// Converts a string to an integer list with "," as delimeter
// No error checking so know what input you are sending
func strToIntList(str string) []int {
	var intList []int
	strList := strings.Split(str, ",")
	for _, str := range strList {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			panic("Failed to convert str to int")
		}
		intList = append(intList, intVal)
	}
	return intList
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

func dayFour() {
	file := getInputData("bingo/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var lottoNumbers []int
	var boards []*bingo.BingoBoard
	// Read first line for bingonumbers
	scanner.Scan()
	lottoNumbers = strToIntList(scanner.Text())
	// Read bingoboards
	board := bingo.NewBingoBoard()
	var position = 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.TrimSpace(text) != "" {
			// Split on spaces, will have sideeffects. Can't rely on range index
			bingoRow := strings.Split(text, " ")
			for _, bingoNumber := range bingoRow {
				var bingoNumberInt int
				var err error
				// Avoid any blankspaces in front of numbers for example _3
				if strings.TrimSpace(bingoNumber) != "" {
					bingoNumberInt, err = strconv.Atoi(bingoNumber)
					if err != nil {
						panic(err.Error())
					}
					board.Board[bingoNumberInt] = &bingo.BoardMeta{
						Position: position,
						Marked:   false,
					}
					position++
				}
			}
		} else if position != 0 {
			boards = append(boards, board)
			board = bingo.NewBingoBoard()
			position = 0
		}
	}
	// Don't miss last board
	boards = append(boards, board)
	task1 := bingo.PlayBingo(boards, lottoNumbers)
	task2 := bingo.CheckLastBingo(boards, lottoNumbers)
	fmt.Printf("Day 4 result A: %d\n", task1)
	fmt.Printf("Day 4 result B: %d\n", task2)
}

func main() {
	//dayOne()
	//dayTwo()
	//dayThree()
	dayFour()
}
