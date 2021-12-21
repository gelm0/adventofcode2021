package smokebasin

import "fmt"
import "sort"

func checkEdges(value int, compared ...int) bool {
	for _, cmp := range compared {
		if cmp <= value {
			return false
		}
	}
	return true
}

func findMinimas(points []int, height int) (risk int, coordinates []int) {
	length := len(points) / height
	edgeMinima := false
	for i, val := range points {
		row := i / length
		col := i % length
		if row == height-1 {
			if col == 0 {
				edgeMinima = checkEdges(val, points[i+1], points[i-length])
			} else if col == length-1 {
				edgeMinima = checkEdges(val, points[i-1], points[i-length])
			} else {
				edgeMinima = checkEdges(val, points[i-1], points[i-length], points[i+1])
			}
		} else if row == 0 {
			if col == 0 {
				edgeMinima = checkEdges(val, points[i+1], points[i+length])
			} else if col == length-1 {
				edgeMinima = checkEdges(val, points[i-1], points[i+length])
			} else {
				edgeMinima = checkEdges(val, points[i-1], points[i+length], points[i+1])
			}
		} else if col == 0 {
			edgeMinima = checkEdges(val, points[i+length], points[i+1], points[i-length])
		} else if col == length-1 {
			edgeMinima = checkEdges(val, points[i-1], points[i+length], points[i-length])
		} else {
			edgeMinima = checkEdges(val, points[i-1], points[i+length], points[i-length], points[i+1])
		}
		if edgeMinima {
			risk += val + 1
			coordinates = append(coordinates, i)
		}
	}
	return risk, coordinates
}

func FindMinimaRisk(points []int, height int) int {
	risk, _ := findMinimas(points, height)
	return risk
}

func printGrid(points []int, visited []bool, length, height int) {
	prevRow := 0
	for i, _ := range points {
		row := i / length
		if row != prevRow {
			prevRow += 1
			fmt.Println()
		}
		fmt.Printf("%d: %t ", points[i], visited[i])
	}
	fmt.Println()
}

func searchBasins(points, toVisit []int, visited []bool, length, height, found int) int {
	if len(toVisit) == 0 {
		return found
	}
	cord := toVisit[0]
	// Shorten list by 1
	toVisit = toVisit[1:]
	// Move next
	left := cord - 1
	right := cord + 1
	top := cord - length
	bottom := cord + length
	// Matrix cords
	row := cord / length
	col := cord % length
	// Update found values
	if !visited[cord] {
		found += 1
		visited[cord] = true
	}
	if col != 0 && points[left] > points[cord] && points[left] != 9 {
		toVisit = append(toVisit, left)
	}
	if col != length-1 && points[right] > points[cord] && points[right] != 9 {
		toVisit = append(toVisit, right)

	}
	if row != 0 && points[top] > points[cord] && points[top] != 9 {
		toVisit = append(toVisit, top)

	}
	if row != height-1 && points[bottom] > points[cord] && points[bottom] != 9 {
		toVisit = append(toVisit, bottom)
	}
	return searchBasins(points, toVisit, visited, length, height, found)
}
func FindBasins(points []int, height int) int {
	length := len(points) / height
	_, minimas := findMinimas(points, height)
	found := 0
	var toVisit []int
	var foundValues []int
	visited := make([]bool, len(points))
	for i, _ := range minimas {
		toVisit = append(toVisit, minimas[i])
		foundValues = append(foundValues, searchBasins(points, toVisit, visited, length, height, found))
	}
	sort.Ints(foundValues)
	return foundValues[len(foundValues)-1] * foundValues[len(foundValues)-2] * foundValues[len(foundValues)-3]
}
