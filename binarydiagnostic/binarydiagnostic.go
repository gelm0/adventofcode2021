package binarydiagnostic

import (
	"strconv"
)

type Rate struct {
	countOne  int
	countZero int
}

func initializeRateMap(mapLen int) map[int]*Rate {
	frequencies := make(map[int]*Rate)
	for i := 0; i < mapLen; i++ {
		frequencies[i] = &Rate{
			countOne:  0,
			countZero: 0,
		}
	}
	return frequencies
}

func getBinaryFrequencies(binaryData []string) (map[int]*Rate, int) {
	// Assume all input data is uniform
	rowLen := len(binaryData[0])
	frequencies := initializeRateMap(rowLen)
	for _, row := range binaryData {
		for i, val := range row {
			if val == '1' {
				frequencies[i].countOne += 1
			} else {
				// Equal to 0
				frequencies[i].countZero += 1
			}
		}
	}
	return frequencies, rowLen
}

func Diagnostic(binaryData []string) int {
	frequencies, rowLen := getBinaryFrequencies(binaryData)
	gamma := 0
	epsilon := 0
	for i := 0; i < rowLen; i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if frequencies[i].countOne < frequencies[i].countZero {
			epsilon = epsilon | 1
		} else {
			gamma = gamma | 1
		}
	}
	return gamma * epsilon
}

// Diagnoses Oxygen and CO2 depending on third input parameter
func DiagnoseOxygen(binaryData []string) int64 {
	return diagnoseGasses(binaryData, 0, true)
}

func DiagnoseCO2(binaryData []string) int64 {
	return diagnoseGasses(binaryData, 0, false)
}

func diagnoseGasses(binaryData []string, place int, measureOxygen bool) int64 {
	var mostImportantBit byte = '0'
	var leastImportantBit byte = '1'
	if measureOxygen {
		mostImportantBit = '1'
		leastImportantBit = '0'
	}
	frequencies, _ := getBinaryFrequencies(binaryData)
	frequency := frequencies[place]
	var gas []string
	for _, binstring := range binaryData {
		if frequency.countOne > frequency.countZero {
			// We want 1s
			if binstring[place] == mostImportantBit {
				gas = append(gas, binstring)
			}
		} else if frequency.countOne == frequency.countZero {
			// We want 1s
			if binstring[place] == mostImportantBit {
				gas = append(gas, binstring)
			}
		} else {
			// 0s are dominant
			if binstring[place] == leastImportantBit {
				gas = append(gas, binstring)
			}
		}
	}
	if len(gas) == 1 {
		i, err := strconv.ParseInt(gas[0], 2, 64)
		if err != nil {
			panic("Can't parse int")
		}
		return i
	}
	place += 1
	return diagnoseGasses(gas, place, measureOxygen)
}
