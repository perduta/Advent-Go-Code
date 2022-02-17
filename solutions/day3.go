package solutions

import (
	"log"
	"math"
)

func Bin2Dec(n []int) int {
	sum := 0
	maxIndex := len(n) - 1
	for i := maxIndex; i >= 0; i-- {
		sum += n[i] * int(math.Pow(2, float64(maxIndex-i)))
	}
	return sum
}

func Day3Part1(inputs []string) int {
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

	log.Println(rowZerosCount)

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

	log.Println(gammaRateBits, epsilonRateBits)
	log.Println(gammaRate, epsilonRate)

	return gammaRate * epsilonRate
}