package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/limwa/advent-of-code-2024/day2/part1"
	"github.com/limwa/advent-of-code-2024/day2/part2"
	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

func solvePart1(input string) string {
	count := 0
	for _, report := range strings.Split(input, "\n") {
		levels := cast.ToIntSlice(strings.Split(report, " ")) 
		if part1.IsSafe(levels) {
			count += 1
		}
	}

	return cast.ToString(count)
}

func solvePart2(input string) string {
	count := 0
	for _, report := range strings.Split(input, "\n") {
		levels := cast.ToIntSlice(strings.Split(report, " ")) 
		if part2.IsSafe(levels) {
			count += 1
		}
	}

	return cast.ToString(count)
}

// START template

//go:embed input.txt
var _input string

func init() {
	util.NormalizeInput(&_input)
}

func measure(execute func() string) string {
    start := time.Now()
	answer := execute()
	fmt.Printf("Execution time: %v\n", time.Since(start))
	return answer
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "which part to solve")
	flag.Parse()

	answer := measure(func() string {
		if part == 1 {
			return solvePart1(_input)
		} else if part == 2 {
			return solvePart2(_input)
		} else {
			panic("a valid part must be specified")
		}
	})
	
	util.CopyToClipboard(answer)
	fmt.Printf("Answer for part %d:\n%s\n", part, answer)
}

// END template
