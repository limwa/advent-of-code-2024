package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
	"github.com/limwa/advent-of-code-2024/lib/math"
)

func solvePart1(input string) string {

	left, right := []int{}, []int{}

	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, "   ")

		first, second := cast.ToInt(numbers[0]), cast.ToInt(numbers[1])

		left = append(left, first)
		right = append(right, second)
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		distance := math.Abs(left[i] - right[i])
		totalDistance += distance
	}

	return cast.ToString(totalDistance)
}

func solvePart2(input string) string {

	leftCounts, rightCounts := map[int]int{}, map[int]int{}

	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Split(line, "   ")

		first, second := cast.ToInt(numbers[0]), cast.ToInt(numbers[1])

		leftCounts[first]++
		rightCounts[second]++
	}

	similarityScore := 0

	for key, count := range leftCounts {
		similarityScore += key * count * rightCounts[key]
	}

	return cast.ToString(similarityScore)
}

// START template

//go:embed input.txt
var _input string

func init() {
	util.NormalizeInput(&_input)
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part to solve")
	flag.Parse()

	var answer string
	if part == 1 {
		answer = solvePart1(_input)
	} else if part == 2 {
		answer = solvePart2(_input)
	} else {
		panic("a valid part must be specified")
	}

	util.CopyToClipboard(answer)
	fmt.Printf("Answer for part %d:\n%s\n", part, answer)
}

// END template