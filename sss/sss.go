package sss

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Six segments
// 0 uses six segments
// 6 uses six segments
// 9 uses six segments

// Five segments
// 2 uses five segments
// 3 uses five segments
// 5 uses five segments

// Unique
// 1 uses two segments
// 4 uses four segments
// 7 uses three segments
// 8 uses seven segments

func UniqueOutput(output []string) int {
	unique := 0
	for _, val := range output {
		strLen := len(val)
		if strLen == 2 || strLen == 4 || strLen == 3 || strLen == 7 {
			unique += 1
		}
	}
	return unique
}

func insertDecodedMap(inputMap map[int][]string, key int, s string) map[int][]string {
	runeList := inputMap[key]
	runeList = append(runeList, s)
	inputMap[key] = runeList
	return inputMap

}

func createInputMap(input []string) map[int][]string {
	inputMap := make(map[int][]string)
	for _, s := range input {
		strLen := len(s)
		switch strLen {
		// 1
		case 2:
			inputMap = insertDecodedMap(inputMap, 1, s)
			// 7
		case 3:
			inputMap = insertDecodedMap(inputMap, 7, s)
			// 4
		case 4:
			inputMap = insertDecodedMap(inputMap, 4, s)
			// 2, 3, 5
		case 5:
			inputMap = insertDecodedMap(inputMap, 2, s)
			inputMap = insertDecodedMap(inputMap, 3, s)
			inputMap = insertDecodedMap(inputMap, 5, s)
			// 0, 6, 9
		case 6:
			inputMap = insertDecodedMap(inputMap, 0, s)
			inputMap = insertDecodedMap(inputMap, 6, s)
			inputMap = insertDecodedMap(inputMap, 9, s)
			// 8
		case 7:
			inputMap = insertDecodedMap(inputMap, 8, s)
		}
	}
	return inputMap
}

// Takes two strings and returns the string which is not contained in the other
// Only returns a string fi they are a difference of 1, otherwise it means that
// the system is non solvable
func stringDifference(list []string, s2 string) (notContained string) {
	for _, v := range list {
		notContained = ""
		for _, r := range v {
			s := string(r)
			if !strings.Contains(s2, s) {
				notContained += s
			}

		}
		if len(notContained) == 1 {
			return notContained
		}
	}
	return ""
}

// Removes all runes in a given input map with value as string list
func removeRune(inputMap map[int][]string, remove string) map[int][]string {
	for k, v := range inputMap {
		for i, s := range v {
			var newString string
			for _, r := range s {
				if string(r) != remove {
					newString += string(r)
				}
			}
			inputMap[k][i] = newString
		}
	}
	return inputMap
}

func printInputMap(inputMap map[int][]string) {
	for k, v := range inputMap {
		fmt.Printf("%d: ", k)
		fmt.Print(v)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func findMap(inputMap map[int][]string, test func(s string) bool) string {
	for _, v := range inputMap {
		for _, s := range v {
			if test(s) {
				return s
			}
		}
	}
	return ""
}

func constructDecodedMap(decodedList []string) map[string]string {
	decodedMap := make(map[string]string)
	fullInput := "abcdefg"
	// zero
	zero := ""
	two := ""
	three := ""
	five := ""
	six := ""
	nine := ""
	for _, r := range fullInput {
		s := string(r)
		// zero
		if s != decodedList[3] {
			zero += s
		}
		// two
		if s != decodedList[1] && s != decodedList[5] {
			two += s
		}
		// three
		if s != decodedList[1] && s != decodedList[4] {
			three += s
		}
		// five
		if s != decodedList[2] && s != decodedList[4] {
			five += s
		}
		// six
		if s != decodedList[2] {
			six += s
		}
		// nine
		if s != decodedList[4] {
			nine += s
		}
	}
	decodedMap[zero] = "0"
	decodedMap[two] = "2"
	decodedMap[three] = "3"
	decodedMap[five] = "5"
	decodedMap[six] = "6"
	decodedMap[nine] = "9"
	return decodedMap
}

func SortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

// Map key val correspondence places
// 	0
// 1 2
//  3
// 4 5
//  6
// We solve the input as a linear equation system
func DecodeInput(input []string, output []string) int {
	inputMap := createInputMap(input)
	decodedList := make([]string, 7)
	// Find position 0 (Top)
	diff71 := stringDifference(inputMap[7], inputMap[1][0])
	inputMap = removeRune(inputMap, diff71)
	// Find position 6 (Bottom)
	diff94 := stringDifference(inputMap[9], inputMap[4][0])
	inputMap = removeRune(inputMap, diff94)
	// Find position 3 (Middle)
	diff31 := stringDifference(inputMap[3], inputMap[1][0])
	inputMap = removeRune(inputMap, diff31)
	// Find position 4 (Left bottom)
	diff84 := stringDifference(inputMap[8], inputMap[4][0])
	inputMap = removeRune(inputMap, diff84)
	// Find position 5 (Right top)
	rem2 := findMap(inputMap, func(s string) bool { return len(s) == 1 })
	inputMap = removeRune(inputMap, rem2)
	// Find position 5 (Right bottom)
	rem1 := inputMap[1][0]
	inputMap = removeRune(inputMap, rem1)
	// Find position 1 (Left top)
	rem8 := inputMap[8][0]
	decodedList[0] = diff71
	decodedList[1] = rem8
	decodedList[2] = rem2
	decodedList[3] = diff31
	decodedList[4] = diff84
	decodedList[5] = rem1
	decodedList[6] = diff94
	outputValue := ""
	decodedMap := constructDecodedMap(decodedList)
	for _, val := range output {
		if len(val) == 2 {
			outputValue += "1"
		} else if len(val) == 3 {
			outputValue += "7"
		} else if len(val) == 4 {
			outputValue += "4"
		} else if len(val) == 7 {
			outputValue += "8"
		} else {
			outputValue += decodedMap[val]
		}
	}
	i, err := strconv.Atoi(outputValue)
	if err != nil {
		panic("Can't convert value")
	}
	return i
}
