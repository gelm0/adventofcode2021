package hydtrothermal

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const ventSize = 1000

func findVentsLargerThenTwo(ventMap [][]int) int {
	countPlaces := 0
	for i := 0; i < ventSize; i++ {
		for j := 0; j < ventSize; j++ {
			if ventMap[i][j] >= 2 {
				countPlaces += 1
			}
		}
	}
	return countPlaces
}

func HorizontalVerticalVents(file *os.File) int {
	scanner := bufio.NewScanner(file)
	ventMap := make([][]int, ventSize)
	for i := 0; i < ventSize; i++ {
		ventMap[i] = make([]int, ventSize)
	}
	for scanner.Scan() {
		text := scanner.Text()
		var x0, y0, x1, y1 int
		fmt.Sscanf(text, "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if x0 == x1 {
			if y1 > y0 {
				for i := y0; i <= y1; i++ {
					ventMap[i][x0] += 1
				}
			} else {
				for i := y1; i <= y0; i++ {
					ventMap[i][x0] += 1
				}
			}
		}
		if y0 == y1 {
			if x1 > x0 {
				for i := x0; i <= x1; i++ {
					ventMap[y0][i] += 1
				}
			} else {
				for i := x1; i <= x0; i++ {
					ventMap[y0][i] += 1
				}
			}
		}
	}
	return findVentsLargerThenTwo(ventMap)
}

func HorizontalVerticalDiagonalVents(file *os.File) int {
	scanner := bufio.NewScanner(file)
	ventMap := make([][]int, ventSize)
	for i := 0; i < ventSize; i++ {
		ventMap[i] = make([]int, ventSize)
	}
	for scanner.Scan() {
		text := scanner.Text()
		var x0, y0, x1, y1 int
		fmt.Sscanf(text, "%d,%d -> %d,%d", &x0, &y0, &x1, &y1)
		if x0 == x1 {
			if y1 > y0 {
				for i := y0; i <= y1; i++ {
					ventMap[i][x0] += 1
				}
			} else {
				for i := y1; i <= y0; i++ {
					ventMap[i][x0] += 1
				}
			}
		}
		if y0 == y1 {
			if x1 > x0 {
				for i := x0; i <= x1; i++ {
					ventMap[y0][i] += 1
				}
			} else {
				for i := x1; i <= x0; i++ {
					ventMap[y0][i] += 1
				}
			}
		}
		// 45 Degrees
		if math.Abs(float64(x0-x1)) == math.Abs(float64(y0-y1)) {
			var x int
			var y int
			var endX int

			if x0 > x1 && y0 > y1 {
				y = y1
				x = x1
				endX = x0
				for x <= endX {
					ventMap[y][x] += 1
					y += 1
					x += 1
				}
			}
			if x0 < x1 && y0 > y1 {
				y = y0
				x = x0
				endX = x1
				for x <= endX {
					ventMap[y][x] += 1
					y -= 1
					x += 1
				}
			}
			if x0 > x1 && y0 < y1 {
				y = y0
				x = x0
				endX = x1
				for x >= endX {
					ventMap[y][x] += 1
					y += 1
					x -= 1
				}
			}
			if x0 < x1 && y0 < y1 {
				y = y0
				x = x0
				endX = x1
				for x <= endX {
					ventMap[y][x] += 1
					y += 1
					x += 1
				}
			}
		}
	}
	return findVentsLargerThenTwo(ventMap)
}
