package main

import (
	_ "embed"
	"flag"
	"fmt"
	"maps"
	"strings"

	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type Point struct {
	X int
	Y int
}

type Map struct {
	Width int
	Height int
	Obstacles map[Point]bool
	GuardPosition Point
	GuardDirection Direction
}

type Direction int

const (
	Up Direction = 0
	Down Direction = 1
	Left Direction = 2
	Right Direction = 3
)

func (m *Map) IsObstacle(p Point) bool {
	_, ok := m.Obstacles[p];
	return ok
}

func (m *Map) PlaceObstacle(p Point) {
	m.Obstacles[p] = true
}

func (m *Map) RemoveObstacle(p Point) {
	delete(m.Obstacles, p)
}

func (m *Map) IsInBounds(p Point) bool {
	return p.X >= 0 && p.X < m.Width && p.Y >= 0 && p.Y < m.Height
}

func (m *Map) GetNextGuardDirection() Direction {
	switch m.GuardDirection {
		case Up:
			return Right
		case Right:
			return Down
		case Down:
			return Left
		case Left:
			return Up
		default:
			panic("invalid guard direction")
	}
}

func (m *Map) Update() {
	nextPosition := Move(m.GuardPosition, m.GuardDirection)
	if m.IsObstacle(nextPosition) {
		m.GuardDirection = m.GetNextGuardDirection()
	} else {
		m.GuardPosition = nextPosition
	}
}

func Move(p Point, d Direction) Point {
	switch d {
		case Up:
			p.Y -= 1
		case Right:
			p.X += 1
		case Down:
			p.Y += 1
		case Left:
			p.X -= 1
		default:
			panic("invalid direction")
	}

	return p
}

func GetGuardDirectionFromInput(character rune) Direction {
	switch character {
		case '^':
			return Up
		case '>':
			return Right
		case 'v':
			return Down
		case '<':
			return Left
		default:
			panic("invalid direction")
	}
}

func CreateMapFromInput(input string) Map {
	lines := strings.Split(input, "\n")

	puzzle := Map{
		Width: len(lines[0]),
		Height: len(lines),
		Obstacles: map[Point]bool{},
		GuardPosition: Point{
			X: -1,
			Y: -1,
		},
		GuardDirection: Up,
	}

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				puzzle.PlaceObstacle(Point{X: x, Y: y})
				continue
			} else if c == '.' {
				continue
			}

			puzzle.GuardPosition = Point{X: x, Y: y}
			puzzle.GuardDirection = GetGuardDirectionFromInput(c)
		}
	}

	return puzzle
}

func solvePart1(input string) string {
	puzzle := CreateMapFromInput(input)
	visitedPositions := map[Point]bool{}

	for puzzle.IsInBounds(puzzle.GuardPosition) {
		visitedPositions[puzzle.GuardPosition] = true
		puzzle.Update()
	}
	
	return cast.ToString(len(visitedPositions))
}

type VisitHistoryEntry struct {
	Pos Point
	Dir Direction
}

func CreateVisitHistoryEntry(m *Map) VisitHistoryEntry {
	return VisitHistoryEntry{
		Pos: m.GuardPosition,
		Dir: m.GuardDirection,
	}
}

func solvePart2(input string) string {
	puzzle := CreateMapFromInput(input)

	history := map[VisitHistoryEntry]bool{}
	visitedPositions := map[Point]bool{}
	
	count := 0
	for puzzle.IsInBounds(puzzle.GuardPosition) {
		history[CreateVisitHistoryEntry(&puzzle)] = true
		visitedPositions[puzzle.GuardPosition] = true

		// Check if placing an obstacle in front of the guard
		// would eventually cause the guard to go back to an
		// already visited position.
		obstaclePosition := Move(puzzle.GuardPosition, puzzle.GuardDirection)
		if puzzle.IsInBounds(obstaclePosition) && !puzzle.IsObstacle(obstaclePosition) && !visitedPositions[obstaclePosition] {
			temporaryHistory := maps.Clone(history)

			initialPosition := puzzle.GuardPosition
			initialDirection := puzzle.GuardDirection

			puzzle.PlaceObstacle(obstaclePosition)
			puzzle.GuardDirection = puzzle.GetNextGuardDirection()

			isLooping := false
			for puzzle.IsInBounds(puzzle.GuardPosition) {
				vhe := CreateVisitHistoryEntry(&puzzle)
				if _, ok := temporaryHistory[vhe]; ok {
					isLooping = true
					break
				}

				temporaryHistory[vhe] = true
				puzzle.Update()
			}

			puzzle.RemoveObstacle(obstaclePosition)
			puzzle.GuardDirection = initialDirection
			puzzle.GuardPosition = initialPosition

			if isLooping {
				count++
			}
		}

		puzzle.Update()
	}

	return cast.ToString(count)
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
