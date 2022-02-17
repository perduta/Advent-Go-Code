package solutions

import (
	"log"
	"math"
	"strconv"
)

func Bin2Dec(n []int) int {
	sum := 0
	maxIndex := len(n) - 1
	for i := maxIndex; i >= 0; i-- {
		sum += n[i] * int(math.Pow(2, float64(maxIndex-i)))
	}
	return sum
}

func GetZerosCount(inputs []string) (int, int, []int) {
	rowLength := len(inputs[0])
	rowZerosCount := make([]int, rowLength)
	reportLength := len(inputs)

	for _, input := range inputs {
		for i, c := range input {
			if c == '0' {
				rowZerosCount[i]++
			}
		}
	}
	return rowLength, reportLength, rowZerosCount
}

func Day3Part1(inputs []string) int {
	rowLength, reportLength, rowZerosCount := GetZerosCount(inputs)

	gammaRateBits := make([]int, rowLength)
	epsilonRateBits := make([]int, rowLength)
	for i := 0; i < rowLength; i++ {
		isZeroMostCommon := rowZerosCount[i] > reportLength/2
		if isZeroMostCommon {
			gammaRateBits[i] = 0
			epsilonRateBits[i] = 1
		} else {
			gammaRateBits[i] = 1
			epsilonRateBits[i] = 0
		}
	}

	gammaRate := Bin2Dec(gammaRateBits)
	epsilonRate := Bin2Dec(epsilonRateBits)

	return gammaRate * epsilonRate
}

func filterByNByte(inputs []string, n int, b byte) (ret []string) {
	for _, input := range inputs {
		if input[n] == b {
			ret = append(ret, input)
		}
	}
	return
}

func Day3Part2Value(inputs []string, more, less, equal byte) int {
	rowLength, _, _ := GetZerosCount(inputs)
	for i := 0; i < rowLength; i++ {
		_, reportLength, rowZerosCount := GetZerosCount(inputs)
		nZerosMore := rowZerosCount[i]  - (reportLength - rowZerosCount[i])
		if nZerosMore > 0 {
			inputs = filterByNByte(inputs, i, more)			
		} else if nZerosMore < 0 {
			inputs = filterByNByte(inputs, i, less)
		} else if nZerosMore == 0 {
			inputs = filterByNByte(inputs, i, equal)
		}
		if len(inputs) <= 1 {
			break
		}
	}

	ret, err := strconv.ParseInt(inputs[0], 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(ret)
}

func Day3Part2(inputs []string) int {
	oxygenGeneratorRating := Day3Part2Value(inputs, '0', '1', '1')
	co2ScrubberRating := Day3Part2Value(inputs, '1', '0', '0')
	return int(oxygenGeneratorRating) * int(co2ScrubberRating)
}