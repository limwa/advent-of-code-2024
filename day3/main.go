package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

func solvePart1(input string) string {

	re := regexp.MustCompile(`mul\((?P<first>\d{1,3}),(?P<second>\d{1,3})\)`)
	
	firstIndex := re.SubexpIndex("first")
	secondIndex := re.SubexpIndex("second")

	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range matches {
		first := cast.ToInt(match[firstIndex])
		second := cast.ToInt(match[secondIndex])

		sum += first * second
	}

	return cast.ToString(sum)
}

func solvePart2(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	activeRe := regexp.MustCompile(`(?:do\(\)|^).*?(?:don't\(\)|$)`);

	sum := 0
	for _, match := range activeRe.FindAllStringSubmatch(input, -1) {
		active := match[0]
		sum += cast.ToInt(solvePart1(active))
	}

	return cast.ToString(sum)
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
