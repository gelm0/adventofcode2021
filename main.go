package main

import (
	"adventofcode/binarydiagnostic"
	"adventofcode/bingo"
	"adventofcode/dive"
	"adventofcode/dumbooctopus"
	hydtrothermal "adventofcode/hydrothermal"
	"adventofcode/lanternfish"
	"adventofcode/passagepathing"
	"adventofcode/smokebasin"
	"adventofcode/sonarsweep"
	"adventofcode/sss"
	"adventofcode/syntaxscoring"
	"adventofcode/transparentorigami"
	"adventofcode/whaletreachery"
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func dayFive() {
	file := getInputData("hydrothermal/input.txt")
	defer file.Close()
	task1 := hydtrothermal.HorizontalVerticalVents(file)
	// Reset file pointer
	file.Seek(0, io.SeekStart)
	task2 := hydtrothermal.HorizontalVerticalDiagonalVents(file)
	fmt.Printf("Day 5 result A: %d\n", task1)
	fmt.Printf("Day 5 result B %d\n", task2)
}

func daySix() {
	file := getInputData("lanternfish/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	scanner.Scan()
	strFishList := strings.Split(scanner.Text(), ",")
	initialFish := make([]int, 9)
	for _, s := range strFishList {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			// Worlds most descriptive error, fish the shell or just fish?
			panicStr := fmt.Sprintf("Failed to convert fishstr to fishint: %s", s)
			panic(panicStr)
		}
		initialFish[intVal] += 1
	}
	task1 := lanternfish.GetNumberOfFish(initialFish, 80)
	task2 := lanternfish.GetNumberOfFish(initialFish, 256)
	fmt.Printf("Day 6 result A: %d\n", task1)
	fmt.Printf("Day 6 result B %d\n", task2)
}

func daySeven() {
	file := getInputData("whaletreachery/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	scanner.Scan()
	crabStrList := strings.Split(scanner.Text(), ",")
	var crabPositions []int
	for _, s := range crabStrList {
		intVal, err := strconv.Atoi(s)
		if err != nil {
			panicStr := fmt.Sprintf("Failed to convert crabstr to crabint: %s", s)
			panic(panicStr)
		}
		crabPositions = append(crabPositions, intVal)
	}
	task1 := whaletreachery.CrabDance(crabPositions)
	task2 := whaletreachery.AdvancedCrabDancing(crabPositions)
	fmt.Printf("Day 7 result A: %d\n", task1)
	fmt.Printf("Day 7 result B %d\n", task2)
}

func dayEight() {
	file := getInputData("sss/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var input, output, totalOutput []string
	task2 := 0
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " ")
		switchToOutput := false
		for _, s := range values {
			if s == "|" {
				switchToOutput = true
				continue
			}
			s = sss.SortString(s)
			if !switchToOutput {
				input = append(input, s)
			} else {
				output = append(output, s)
			}
		}
		task2 += sss.DecodeInput(input, output)
		input = []string{}
		totalOutput = append(totalOutput, output...)
		output = []string{}
	}
	task1 := sss.UniqueOutput(totalOutput)
	fmt.Printf("Day 8 result A: %d\n", task1)
	fmt.Printf("Day 8 result B: %d\n", task2)
}

func dayNine() {
	file := getInputData("smokebasin/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	height := 0
	var points []int
	for scanner.Scan() {
		height += 1
		for _, s := range scanner.Text() {
			i := int(s - '0')
			points = append(points, i)
		}
	}
	task1 := smokebasin.FindMinimaRisk(points, height)
	task2 := smokebasin.FindBasins(points, height)
	fmt.Printf("Day 9 result A: %d\n", task1)
	fmt.Printf("Day 9 result B: %d\n", task2)
}

func dayTen() {
	file := getInputData("syntaxscoring/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	task1 := 0
	var autoCompletedScore []int
	for scanner.Scan() {
		score, toComplete := syntaxscoring.ScoreSyntax(scanner.Text())
		task1 += score
		if toComplete != nil {
			autoCmpScore := syntaxscoring.ScoreAutoComplete(toComplete)
			autoCompletedScore = append(autoCompletedScore, autoCmpScore)
		}
	}
	sort.Ints(autoCompletedScore)
	middleScore := len(autoCompletedScore) / 2
	task2 := autoCompletedScore[middleScore]
	fmt.Printf("Day 10 result A: %d\n", task1)
	fmt.Printf("Day 10 result B: %d\n", task2)
}

func dayEleven() {
	file := getInputData("dumbooctopus/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var points []int
	height := 0
	for scanner.Scan() {
		height += 1
		text := scanner.Text()
		for _, r := range text {
			i := int(r - '0')
			points = append(points, i)
		}
	}
	nodes := dumbooctopus.InitializeNodes(points, height)
	task1 := dumbooctopus.Step(100, nodes)
	task2 := dumbooctopus.FindFirstSynchronized(nodes)
	fmt.Printf("Day 11 result A: %d\n", task1)
	fmt.Printf("Day 11 result B: %d\n", task2+100)
}

func dayTwelve() {
	file := getInputData("passagepathing/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	var nodes []string
	graph := passagepathing.NewGraph()
	for scanner.Scan() {
		nodes = strings.Split(scanner.Text(), "-")
		graph.AddNode(nodes[0])
		graph.AddNode(nodes[1])
		graph.AddEdge(nodes[0], nodes[1])
		graph.AddEdge(nodes[1], nodes[0])
	}
	smallCaveCounter := make(map[string]int)

	paths := []string{}
	task1 := 0
	task2 := 0
	graph.PathsWithOneSmallCave("start", "end", &task1)
	graph.PathsWithTwoSmallCaves("start", "end", paths, smallCaveCounter, &task2)
	fmt.Printf("Day 12 result A: %d\n", task1)
	fmt.Printf("Day 12 result B: %d\n", task2)
}

func dayThirteen() {
	file := getInputData("transparentorigami/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	parseFolds := false
	var points []int
	foldRune := []rune{}
	foldPos := []int{}
	var maxX, maxY int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			parseFolds = true
		}
		// Parse points
		if !parseFolds {
			xy := strings.Split(text, ",")
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				panic("Failed to convert x")
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				panic("Failed to convert y")
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			points = append(points, x)
			points = append(points, y)
			// Parse folds instead
		} else {
			var r rune
			var d int
			fmt.Sscanf(text, "fold along %c=%d", &r, &d)
			foldRune = append(foldRune, r)
			foldPos = append(foldPos, d)
		}
	}
	foldPos = foldPos[1:]
	foldRune = foldRune[1:]

	// Coordinate system ranges from 0 so we have to add 1 to correct
	maxX += 1
	maxY += 1
	firstFold := true
	paper := transparentorigami.MakePaper(maxX, maxY)
	transparentorigami.FillPaper(points, paper)
	for i := 0; i < len(foldPos); i++ {

		if foldRune[i] == 'x' {
			paper = transparentorigami.FoldX(paper, foldPos[i])
		} else {
			paper = transparentorigami.FoldY(paper, foldPos[i])
		}
		if firstFold {
			task1 := transparentorigami.CountPoints(paper)
			fmt.Printf("Day 13 result A: %d\n", task1)
			firstFold = false
		}
	}
	transparentorigami.PrintPaper(paper)
}

func main() {
	// Write a commandline interpreter or use cobra
	//dayOne()
	//dayTwo()
	//dayThree()
	//dayFour()
	//dayFive()
	//daySix()
	//daySeven()
	//dayEight()
	//dayNine()
	//dayTen()
	//dayEleven()
	//dayTwelve()
	dayThirteen()
}
