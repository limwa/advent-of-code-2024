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

type CompactEntry struct {
	Id int
	Start int
	Size int
}

func parseCompactEntries(input string) []CompactEntry {
	digits := cast.ToIntSlice(strings.Split(input, ""))

	entries := []CompactEntry{}
	
	isEmpty := true
	currentPos := 0

	for i := 0; i < len(digits); i++ {
		isEmpty = !isEmpty
		
		blockSize := digits[i]
		entry := CompactEntry{Start: currentPos, Size: blockSize}

		if isEmpty {
			entry.Id = -1
		} else {
			entry.Id = i / 2
		}

		entries = append(entries, entry)
		currentPos += blockSize
	}

	return entries
}

func calculateChecksum(entries []CompactEntry) int {
	checksum := 0

	for _, entry := range entries {
		if entry.Id == -1 {
			continue
		}

		start := entry.Start
		end := entry.Start + entry.Size

		indicesSum := (end - start) * (end + start - 1) / 2
		checksum += indicesSum * entry.Id
	}

	return checksum
}

func solvePart1(input string) string {
	entries := parseCompactEntries(input)

	result := []CompactEntry{}

	currentFillIndex := len(entries) - 1
	currentFillIndex -= currentFillIndex % 2

	for i := 0; i <= currentFillIndex; i++ {
		entry := entries[i]

		if entry.Id >= 0 {
			result = append(result, entry)
			continue
		}

		for entry.Size > 0 && i <= currentFillIndex {
			currentFillEntry := entries[currentFillIndex]

			if currentFillEntry.Size == 0 {
				currentFillIndex -= 2
				continue
			}

			filledSpaces := min(currentFillEntry.Size, entry.Size)

			result = append(
				result,
				CompactEntry{
					Start: entry.Start,
					Size:  filledSpaces,
					Id: currentFillEntry.Id,
				},
			)

			entry.Start += filledSpaces
			entry.Size -= filledSpaces

			currentFillEntry.Size -= filledSpaces
			entries[currentFillIndex] = currentFillEntry

			afterCurrentFillIndex := currentFillIndex + 1
			if afterCurrentFillIndex < len(entries) {
				afterCurrentFillEntry := entries[afterCurrentFillIndex]
				afterCurrentFillEntry.Start -= filledSpaces
				afterCurrentFillEntry.Size += filledSpaces
				entries[afterCurrentFillIndex] = afterCurrentFillEntry
			}
		}
	}

	return cast.ToString(calculateChecksum(result))
}

func solvePart2(input string) string {
	entries := parseCompactEntries(input)

	emptyEntries := []CompactEntry{}
	for i := 1; i < len(entries); i += 2 {
		entry := entries[i]
		if entry.Size == 0 {
			continue
		}

		emptyEntries = append(emptyEntries, entry)
	}

	result := []CompactEntry{}

	lastBlockIndex := len(entries) - 1
	lastBlockIndex -= lastBlockIndex % 2

	for i := lastBlockIndex; i >= 0; i -= 2 {
		entry := entries[i]

		replacementEntry := entry
		
		for j, emptyEntry := range emptyEntries {
			if replacementEntry.Start < emptyEntry.Start {
				break
			}

			if emptyEntry.Size >= replacementEntry.Size {
				replacementEntry = CompactEntry{
					Start: emptyEntry.Start,
					Size: replacementEntry.Size,
					Id: replacementEntry.Id,
				}

				emptyEntry.Size -= replacementEntry.Size
				emptyEntry.Start += replacementEntry.Size

				if emptyEntry.Size == 0 {
					emptyEntries = append(emptyEntries[:j], emptyEntries[j+1:]...)
				} else {
					emptyEntries[j] = emptyEntry
				}
			}
		}

		result = append(result, replacementEntry)
	}

	return cast.ToString(calculateChecksum(result))
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
