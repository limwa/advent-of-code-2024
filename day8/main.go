package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/spatial"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type Frequency struct {
	Antenas []spatial.Vec2D
}

type Puzzle struct {
	Bounds spatial.Vec2D
	Frequencies map[string]Frequency
}

func createPuzzleFromInput(input string) Puzzle {
	lines := strings.Split(input, "\n")
	
	puzzle := Puzzle{
		Bounds: spatial.Vec2D{X: len(lines[0]), Y: len(lines)},
		Frequencies: map[string]Frequency{},
	}

	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}

			frequency, ok := puzzle.Frequencies[string(char)]
			if !ok {
				frequency = Frequency{}
			}
			
			frequency.Antenas = append(frequency.Antenas, spatial.Vec2D{X: x, Y: y})
			puzzle.Frequencies[string(char)] = frequency
		}
	}

	return puzzle
}

func solvePart1(input string) string {
	origin := spatial.Vec2D{X: 0, Y: 0}
	puzzle := createPuzzleFromInput(input)

	unique_antenas := map[spatial.Vec2D]bool{}

	for _, frequency := range puzzle.Frequencies {

		for i := 0; i < len(frequency.Antenas) - 1; i++ {
			firstAntena := frequency.Antenas[i]

			for j := i + 1; j < len(frequency.Antenas); j++ {
				secondAntena := frequency.Antenas[j]

				difference := firstAntena.Sub(secondAntena)

				firstAntinode := firstAntena.Add(difference)
				if firstAntinode.IsWithinBounds(origin, puzzle.Bounds) {
					unique_antenas[firstAntinode] = true
				}

				secondAntinode := secondAntena.Sub(difference)
				if secondAntinode.IsWithinBounds(origin, puzzle.Bounds) {
					unique_antenas[secondAntinode] = true
				}
			}
		}
	}

	return cast.ToString(len(unique_antenas))
}

func solvePart2(input string) string {
	origin := spatial.Vec2D{X: 0, Y: 0}
	puzzle := createPuzzleFromInput(input)

	unique_antenas := map[spatial.Vec2D]bool{}

	for _, frequency := range puzzle.Frequencies {

		for i := 0; i < len(frequency.Antenas) - 1; i++ {
			firstAntena := frequency.Antenas[i]
			
			for j := i + 1; j < len(frequency.Antenas); j++ {
				secondAntena := frequency.Antenas[j]

				difference := firstAntena.Sub(secondAntena)

				currentPoint := firstAntena
				for currentPoint.IsWithinBounds(origin, puzzle.Bounds) {
					unique_antenas[currentPoint] = true
					currentPoint = currentPoint.Add(difference)
				}

				currentPoint = secondAntena
				for currentPoint.IsWithinBounds(origin, puzzle.Bounds) {
					unique_antenas[currentPoint] = true
					currentPoint = currentPoint.Sub(difference)
				}
			}
		}
	}

	return cast.ToString(len(unique_antenas))
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
