package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/limwa/advent-of-code-2024/day10/pathfinding"
	"github.com/limwa/advent-of-code-2024/lib/cast"
	"github.com/limwa/advent-of-code-2024/lib/spatial"
	"github.com/limwa/advent-of-code-2024/lib/util"
)

type Puzzle struct {
	Size spatial.Vec2D
	Grid [][]int
}

func (p *Puzzle) ValueAt(position spatial.Vec2D) int {
	return p.Grid[position.Y][position.X]
}

func parsePuzzle(input string) Puzzle {
	lines := strings.Split(input, "\n")
	size := spatial.Vec2D{X: len(lines[0]), Y: len(lines)}

	grid := make([][]int, size.Y)
	for i, line := range lines {
		grid[i] = cast.ToIntSlice(strings.Split(line, ""))
	}

	return Puzzle{Size: size, Grid: grid}
}

type HikingResult [][]*HikingEntry

type HikingEntry struct {
	Visited bool
	Parents map[spatial.Vec2D]int
}

func createHikingResult(puzzle Puzzle) HikingResult {
	result := make(HikingResult, puzzle.Size.Y)
	for y := range result {
		result[y] = make([]*HikingEntry, puzzle.Size.X)
		for x := range result[y] {
			result[y][x] = &HikingEntry{
				Parents: map[spatial.Vec2D]int{},
			}
		}
	}
	return result
}

func (h *HikingResult) EntryAt(position spatial.Vec2D) *HikingEntry {
	return (*h)[position.Y][position.X]
}

func hikeAndGetResult(input string) (HikingResult, map[spatial.Vec2D]bool) {
	puzzle := parsePuzzle(input)
	result := createHikingResult(puzzle)

	path := pathfinding.PathfindingHeap{}
	heap.Init(&path)

	for y, line := range puzzle.Grid {
		for x, cell := range line {
			if cell == 9 {
				position := spatial.Vec2D{X: x, Y: y}
				heap.Push(&path, &pathfinding.PathfindingItem{Position: position, Value: 9})

				neighborEntry := result.EntryAt(position)
				neighborEntry.Parents[position] = 1
			}
		}
	}

	origin := spatial.Vec2D{X: 0, Y: 0}
	neighbors := []spatial.Vec2D{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
	}

	starts := map[spatial.Vec2D]bool{}

	for path.Len() > 0 {
		item := heap.Pop(&path).(*pathfinding.PathfindingItem)
		
		itemEntry := result.EntryAt(item.Position)
		if itemEntry.Visited {
			continue
		}
		
		itemEntry.Visited = true
		
		if item.Value == 0 {
			starts[item.Position] = true
			continue
		}

		nextValue := item.Value - 1
		for _, neighbor := range neighbors {
			neighborPosition := item.Position.Add(neighbor)
			
			if neighborPosition.IsWithinBounds(origin, puzzle.Size) && puzzle.ValueAt(neighborPosition) == nextValue {
				heap.Push(&path, &pathfinding.PathfindingItem{Position: neighborPosition, Value: nextValue})
				
				neighborEntry := result.EntryAt(neighborPosition)
				for parent, count := range itemEntry.Parents {
					neighborEntry.Parents[parent] += count
				}
			}
		}
	}

	return result, starts
}

func solvePart1(input string) string {
	result, starts := hikeAndGetResult(input)

	sum := 0
	for k := range starts {
		seen := map[spatial.Vec2D]bool{}

		entry := result.EntryAt(k)
		for end := range entry.Parents {
			seen[end] = true
		}

		sum += len(seen)
	}

	return cast.ToString(sum)
}

func solvePart2(input string) string {
	result, starts := hikeAndGetResult(input)

	sum := 0
	for k := range starts {
		entry := result.EntryAt(k)

		for _, count := range entry.Parents {
			sum += count
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
