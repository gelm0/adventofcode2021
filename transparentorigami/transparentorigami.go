package transparentorigami

import "fmt"

// Create an empty sheet
func MakePaper(row, col int) [][]bool {
	paper := make([][]bool, col)
	for i := 0; i < col; i++ {
		paper[i] = make([]bool, row)
	}
	return paper
}

// Fill sheet with points
// Input is a list of len(points) mod 2 == 0,
// where x,y = point[n], point[n+1]
func FillPaper(points []int, paper [][]bool) {
	for i := 0; i < len(points); i += 2 {
		x := points[i]
		y := points[i+1]
		paper[y][x] = true
	}
}

func PrintPaper(paper [][]bool) {
	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			if paper[i][j] {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func FoldX(paper [][]bool, fold int) [][]bool {
	col := len(paper)
	cutPaper := MakePaper(fold, col)
	for i := 0; i < col; i++ {
		count := 1
		for j := fold; j < len(paper[0])-1; j++ {
			orgJ := j - fold
			translatedX := fold - count
			x := j + 1
			if paper[i][x] {
				cutPaper[i][translatedX] = true
			}
			// Restore original points since we are slicing the paper
			if orgJ < fold && paper[i][orgJ] {
				cutPaper[i][orgJ] = true
			}
			count++
		}
	}
	return cutPaper
}

func FoldY(paper [][]bool, fold int) [][]bool {
	col := len(paper)
	cutPaper := MakePaper(len(paper[0]), fold)
	count := 1
	for i := fold; i < col-1; i++ {
		translatedY := fold - count
		orgI := i - fold
		y := i + 1
		for j := 0; j < len(paper[0]); j++ {
			if paper[y][j] {
				cutPaper[translatedY][j] = true
			}
			// Restore original points since we are slicing the paper
			if orgI < fold && paper[orgI][j] {
				cutPaper[orgI][j] = true
			}
		}
		count++
	}
	return cutPaper
}

func CountPoints(paper [][]bool) int {
	count := 0
	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			if paper[i][j] {
				count++
			}
		}
	}
	return count
}
