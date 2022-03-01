package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const SEGMENTS_LEN = 7
const DIGITS = 10

var DIGITS_CONFIGURATION = [][]int{
	{1, 1, 1, 0, 1, 1, 1}, // 0
	{0, 0, 1, 0, 0, 1, 0}, // 1
	{1, 0, 1, 1, 1, 0, 1}, // 2
	{1, 0, 1, 1, 0, 1, 1}, // 3
	{0, 1, 1, 1, 0, 1, 0}, // 4
	{1, 1, 0, 1, 0, 1, 1}, // 5
	{1, 1, 0, 1, 1, 1, 1}, // 6
	{1, 0, 1, 0, 0, 1, 0}, // 7
	{1, 1, 1, 1, 1, 1, 1}, // 8
	{1, 1, 1, 1, 0, 1, 1}, // 9
}

func Day8Part1(inputs []string) int {
	count := 0
	for _, input := range inputs {
		outputValues := strings.Split(input, " | ")[1]
		words := strings.Split(outputValues, " ")
		for _, word := range words {
			if len(word) == 2 || len(word) == 3 || len(word) == 4 || len(word) == 7 {
				count++
			}
		}
	}
	return count
}

func heapPermutation(a []int, size int, ch chan []int) {
	var generate func(a []int, size int, ch chan []int)
	generate = func(a []int, size int, ch chan []int) {
		if size == 1 {
			b := make([]int, len(a))
			copy(b, a)
			ch <- b
			return
		}
		for i := 0; i < size; i++ {
			generate(a, size-1, ch)
			if size%2 == 1 {
				a[0], a[size-1] = a[size-1], a[0]
			} else {
				a[i], a[size-1] = a[size-1], a[i]
			}
		}
	}
	generate(a, size, ch)
	close(ch)
}

func getSegments(word string, permutation []int) []int {
	segments := make([]int, SEGMENTS_LEN)
	for _, segment := range word {
		segments[permutation[segment-'a']] = 1
	}
	return segments
}

func getDigit(segments []int) (int, error) {
	if len(segments) != SEGMENTS_LEN {
		return 0, fmt.Errorf("invalid segments length: %d", len(segments))
	}
	for digit := 0; digit < DIGITS; digit++ {
		if isValidDigit(segments, digit) {
			return digit, nil
		}
	}
	return 0, fmt.Errorf("no valid digit found")

}

func isValidDigit(segments []int, digit int) bool {
	for i := 0; i < len(segments); i++ {
		if segments[i] != DIGITS_CONFIGURATION[digit][i] {
			return false
		}
	}
	return true
}

func joinInts(ints []int) (s string) {
	for _, i := range ints {
		s += strconv.Itoa(i)
	}
	return
}

func decode(words []string, permutation []int) (int, error) {
	decoded := []int{}
	for _, word := range words {
		activeSegments := getSegments(word, permutation)
		digit, err := getDigit(activeSegments)
		if err != nil {
			return 0, err
		}
		decoded = append(decoded, digit)
	}
	ret, err := strconv.Atoi(joinInts(decoded))
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func Brute(input string) int {
	permutations := make(chan []int)
	go heapPermutation([]int{0, 1, 2, 3, 4, 5, 6}, SEGMENTS_LEN, permutations)

	parts := strings.Split(input, " | ")
	seed := strings.Split(parts[0], " ")
	passcode := strings.Split(parts[1], " ")
	for p := range permutations {
		_, err := decode(append(seed, passcode...), p)
		if err != nil {
			continue
		}
		ret, err := decode(passcode, p)
		if err == nil {
			return ret
		}
	}
	return 0
}

func Day8Part2(inputs []string) int {
	sum := 0
	answers := make(chan int, len(inputs))
	var wg sync.WaitGroup
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(s string) {
			defer wg.Done()
			answers <- Brute(s)
		}(input)
	}
	wg.Wait()
	close(answers)
	for answer := range answers {
		sum += answer
	}
	return sum
}
