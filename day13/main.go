package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"time"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/spatial"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type Machine struct {
	ButtonA spatial.Vec2D
	ButtonB spatial.Vec2D
	Prize  spatial.Vec2D
}

type Machines []Machine

func ParseMachines(input string) Machines {
	re := regexp.MustCompile(`Button A: X\+(?<xa>\d+), Y\+(?<ya>\d+)\nButton B: X\+(?<xb>\d+), Y\+(?<yb>\d+)\nPrize: X=(?<xp>\d+), Y=(?<yp>\d+)`)
	
	matches := re.FindAllStringSubmatch(input, -1)
	machines := Machines{}

	for _, match := range matches {
		xa := cast.ToInt(match[re.SubexpIndex("xa")])
		ya := cast.ToInt(match[re.SubexpIndex("ya")])
		xb := cast.ToInt(match[re.SubexpIndex("xb")])
		yb := cast.ToInt(match[re.SubexpIndex("yb")])
		xp := cast.ToInt(match[re.SubexpIndex("xp")])
		yp := cast.ToInt(match[re.SubexpIndex("yp")])

		machine := Machine{
			ButtonA: spatial.Vec2D{X: xa, Y: ya},
			ButtonB: spatial.Vec2D{X: xb, Y: yb},
			Prize: spatial.Vec2D{X: xp, Y: yp},
		}

		machines = append(machines, machine)
	}

	return machines
}

func FindTokensForMachine(machine Machine) (bool, int) {
	// There is an explanation for the following equations in explaination.jpg

	xa := machine.ButtonA.X
	ya := machine.ButtonA.Y
	xb := machine.ButtonB.X
	yb := machine.ButtonB.Y
	xp := machine.Prize.X
	yp := machine.Prize.Y

	validB, b := ExactDiv(yp * xa - xp * ya, yb * xa - xb * ya)
	validA, a := ExactDiv(xp - b * xb, xa)

	return validA && validB, a * 3 + b
}

func ExactDiv(a, b int) (bool, int) {
	return a%b == 0, a/b
}

func solvePart1(input string) string {
	machines := ParseMachines(input)

	tokens := 0
	for _, machine := range machines {
		found, machineTokens := FindTokensForMachine(machine)
		if found {
			tokens += machineTokens
		}
	}

	return cast.ToString(tokens)
}

func solvePart2(input string) string {
	machines := ParseMachines(input)

	tokens := 0
	for _, machine := range machines {
		machine.Prize.X += 10000000000000
		machine.Prize.Y += 10000000000000

		found, machineTokens := FindTokensForMachine(machine)
		if found {
			tokens += machineTokens
		}
	}

	return cast.ToString(tokens)
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
