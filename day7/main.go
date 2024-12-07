package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/lists"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

func solvePart1(input string) string {
	equations := strings.Split(input, "\n")

	sum := 0
	for _, equation := range equations {
		equation_parts := strings.Split(equation, ": ")

		target := cast.ToInt(equation_parts[0])
		values := cast.ToIntSlice(strings.Split(equation_parts[1], " "))

		initialValue := values[0]
		queue := []int{initialValue}

		for i := 1; i < len(values); i++ {
			new_queue := []int{}
			
			for _, value := range queue {
				new_queue = append(new_queue, value + values[i], value * values[i])
			}

			queue = new_queue
		}

		if lists.Contains(queue, target) {
			sum += target
		}
	}
		
	return cast.ToString(sum)
}

func solvePart2(input string) string {
	equations := strings.Split(input, "\n")

	sum := 0
	for _, equation := range equations {
		equation_parts := strings.Split(equation, ": ")

		target := cast.ToInt(equation_parts[0])
		values := cast.ToIntSlice(strings.Split(equation_parts[1], " "))

		initialValue := values[0]
		queue := []int{initialValue}

		for i := 1; i < len(values); i++ {
			new_queue := []int{}
			
			right_value := values[i]
			right_value_length := len(cast.ToString(right_value))

			shift_amount := int(math.Pow10(right_value_length))

			for _, value := range queue {
				new_queue = append(
					new_queue,
					value + right_value,
					value * right_value,
					value * shift_amount + right_value,
				)
			}

			queue = new_queue
		}

		if lists.Contains(queue, target) {
			sum += target
		}
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
