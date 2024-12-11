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

type Puzzle struct {
	width int
	height int
	characters []string
}

type PuzzleMatch struct {
	x int
	y int
	dx int
	dy int
}

func createPuzzle(input string) Puzzle {
	characters := strings.Split(input, "\n")
	height := len(characters)
	width := len(characters[0])

	return Puzzle{
		width,
		height,
		characters,
	}
}

func isInPuzzle(x int, y int, puzzle *Puzzle) bool {
	return x >= 0 && x < puzzle.width && y >= 0 && y < puzzle.height
}

func lookFor(x int, y int, word string, puzzle *Puzzle) []PuzzleMatch {
	firstCharacter := word[0]
	if !isInPuzzle(x, y, puzzle) || puzzle.characters[y][x] != firstCharacter {
		return []PuzzleMatch{}
	}

	matches := []PuzzleMatch{}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			found := true

			nx := x
			ny := y
			
			for i := 1; i < len(word); i++ {
				nx += dx
				ny += dy

				if !isInPuzzle(nx, ny, puzzle) || puzzle.characters[ny][nx] != word[i] {
					found = false
					break
				}
			}

			if found {
				matches = append(matches, PuzzleMatch{
					x: x,
					y: y,
					dx: dx,
					dy: dy,
				})
			}
		}
	}

	return matches
}

func solvePart1(input string) string {
	puzzle := createPuzzle(input)

	count := 0
	for y, row := range puzzle.characters {
		for x := range row {
			matches := lookFor(x, y, "XMAS", &puzzle)
			count += len(matches)
		}
	}

	return cast.ToString(count)
}

func solvePart2(input string) string {
	puzzle := createPuzzle(input)
	
	matches := map[PuzzleMatch]bool{} 
	for y, row := range puzzle.characters {
		for x := range row {
			thisMatches := lookFor(x, y, "MAS", &puzzle)
			for _, match := range thisMatches {
				if match.dx == 0 || match.dy == 0 {
					// Horizontal crosses don't count
					// So we save space by skipping them
					continue
				}

				matches[match] = true
			}
		}
	}

	count := 0
	for match := range matches {
		correspondingMatches := []PuzzleMatch{
			{ x: match.x + match.dx * 2, y: match.y, dx: -match.dx, dy: match.dy },
			{ x: match.x, y: match.y + match.dy * 2, dx: match.dx, dy: -match.dy },
		}

		for _, correspondingMatch := range correspondingMatches {
			if matches[correspondingMatch] {
				count++
			}
		}
	}

	// Crosses are counted twice
	count /= 2

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
