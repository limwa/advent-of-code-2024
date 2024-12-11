package main

import (
	_ "embed"
	"flag"
	"fmt"
	"time"

	"github.com/limwa/advent-of-code-2024/lib/util"
)

func solvePart1(input string) string {
	// TODO: Solve part 1
	return "<TODO>"
}

func solvePart2(input string) string {
	// TODO: Solve part 2
	return "<TODO>"
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
