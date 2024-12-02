package main

import (
	_ "embed"
	"flag"
	"fmt"

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
