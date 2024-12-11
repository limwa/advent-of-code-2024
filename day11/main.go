package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type NumberAtDepth struct {
	Number int
	Depth int
}

func countStones(n, depth int, memo *map[NumberAtDepth]int) int {
	// fmt.Println("countStones", n, depth)
	if depth <= 0 {
		return 1
	}

	entry := NumberAtDepth{Number: n, Depth: depth}

	if memoResult, ok := (*memo)[entry]; ok {
		return memoResult
	}

	var result int
	if n == 0 {
		result = countStones(1, depth - 1, memo)
	} else {
		nstr := cast.ToString(n)
		nlen := len(nstr)

		if nlen % 2 == 0 {
			firstHalf := cast.ToInt(nstr[:nlen/2])
			secondHalf := cast.ToInt(nstr[nlen/2:])

			result = countStones(firstHalf, depth - 1, memo) + countStones(secondHalf, depth - 1, memo)
		} else {
			result = countStones(n * 2024, depth - 1, memo)
		}
	}

	(*memo)[entry] = result
	return result
}

func blink(n int, input string) int {
	numbers := cast.ToIntSlice(strings.Split(input, " "))
	memo := map[NumberAtDepth]int{}

	result := 0
	for _, number := range numbers {
		result += countStones(number, n, &memo)
	}

	return result
}

func solvePart1(input string) string {
	return cast.ToString(blink(25, input))
}

func solvePart2(input string) string {
	return cast.ToString(blink(75, input))
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
