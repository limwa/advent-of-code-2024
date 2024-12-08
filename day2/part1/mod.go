package part1

import (
	"github.com/limwa/advent-of-code-2024/lib/math"
)

func IsSafe(levels []int) bool {
	previous := levels[0]

	isDecreasing := false
	isIncreasing := false

	for i := 1; i < len(levels); i++ {
		current := levels[i]
		diff := current - previous
		previous = current

		if diff < 0 {
			isDecreasing = true
		} else if diff > 0 {
			isIncreasing = true
		} else {
			return false
		}

		absDiff := math.Abs(diff)

		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return !(isDecreasing && isIncreasing)
}
