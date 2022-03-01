package solutions

import (
	"fmt"
	"strings"
	"testing"
)

func day8GetInputs() []string {
	return strings.Split(`be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`, "\n")
}

func TestDay8Part1(t *testing.T) {
	inputs := day8GetInputs()
	expected := 26
	fmt.Println(">>> " + t.Name())
	result := Day8Part1(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay8Part2(t *testing.T) {
	inputs := day8GetInputs()
	expected := 61229
	fmt.Println(">>> " + t.Name())
	result := Day8Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestDay8Part2Full(t *testing.T) {
	inputs := day8GetInputs()
	expected := 61229
	fmt.Println(">>> " + t.Name())
	result := Day8Part2(inputs)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestHeapPermutation(t *testing.T) {
	inputs := []int{1, 2, 3}
	expected := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
		{3, 2, 1},
	}
	fmt.Println(">>> " + t.Name())
	ch := make(chan []int)
	go heapPermutation(inputs, 3, ch)
	i := 0
	ok := 0
	for result := range ch {
		same := true
		for j := 0; j < len(result); j++ {
			if result[j] != expected[i][j] {
				same = false
				break
			}
		}
		if same {
			ok++
		}
		i++
	}
	if ok != len(expected) {
		t.Errorf("Only %d of %d permutations are the same", ok, len(expected))
	}
}

func TestBrute(t *testing.T) {
	inputs := day8GetInputs()
	expected := []int{
		8394,
		9781,
		1197,
		9361,
		4873,
		8418,
		4548,
		1625,
		8717,
		4315,
	}
	for i := 0; i < len(inputs); i++ {
		t.Run(fmt.Sprintf("input %d", i), func(t *testing.T) {
			result := Brute(inputs[i])
			if result != expected[i] {
				t.Errorf("In input #%d expected %d, got %d", i, expected[i], result)
			}
		})
	}

}

func TestDecode(t *testing.T) {
	type args struct {
		sentence    []string
		permutation []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"no valid digit found", args{[]string{""}, []int{0, 1, 2, 3, 4, 5, 6}}, 0, true},
		{"no valid digit found", args{[]string{"a"}, []int{0, 1, 2, 3, 4, 5, 6}}, 0, true},
		{"simple permutation #1", args{[]string{"cf"}, []int{0, 1, 2, 3, 4, 5, 6}}, 1, false},
		{"simple permutation #7", args{[]string{"fca"}, []int{0, 1, 2, 3, 4, 5, 6}}, 7, false},
		{"example", args{strings.Split("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab cdfeb fcadb cdfeb cdbaf", " "), []int{2, 5, 6, 0, 1, 3, 4}}, 85237964015353, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode(tt.args.sentence, tt.args.permutation)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
