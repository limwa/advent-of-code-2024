package part2

func IsSafe(levels []int) bool {
	diffs := make([]int, len(levels) - 1)

	for i := 0; i < len(diffs); i++ {
		diffs[i] = levels[i + 1] - levels[i]
	}

	ensurePositiveDiffs(diffs)
	return areDiffsSafe(diffs)
}

func ensurePositiveDiffs(diffs []int) {
	positives := 0
	negatives := 0

	for _, diff := range diffs {
		if diff > 0 {
			positives += 1
		} else if diff < 0 {
			negatives += 1
		}
	}

	// Ensure positives are always the majority
	if positives < negatives {
		for i := 0; i < len(diffs); i++ {
			diffs[i] = -diffs[i]
		}
	}
}

func areDiffsSafe(diffs []int) bool {
	badIndex := -1
	consecutiveBadIndices := false

	for i := 0; i < len(diffs); i++ {
		if !isDiffSafe(diffs[i]) {
			if badIndex == -1 {
				badIndex = i
			} else if i == badIndex + 1 {
				consecutiveBadIndices = true
			} else {
				return false
			}
		}
	}

	if badIndex == -1 {
		return true
	}

	return (!consecutiveBadIndices && isJoinedDiffSafe(badIndex - 1, badIndex, diffs)) || isJoinedDiffSafe(badIndex, badIndex + 1, diffs)
}

func isDiffSafe(diff int) bool {
	return diff >= 1 && diff <= 3
}

func isJoinedDiffSafe(leftIndex int, rightIndex int, diffs []int) bool {
	// If the badIndex is at an edge, it's always safe to remove it
	if leftIndex < 0 || rightIndex >= len(diffs) {
		return true
	}

	leftDiff := diffs[leftIndex]
	rightDiff := diffs[rightIndex]

	joinedDiff := leftDiff + rightDiff
	return isDiffSafe(joinedDiff)
}

