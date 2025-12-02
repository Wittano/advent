package advent

import (
	"math"
	"slices"
)

type NumberPair struct {
	left, right int
}

func distanceBetweenSmallestNumbers(in []NumberPair) (result int) {
	left, right := make([]int, len(in)), make([]int, len(in))

	for i, pair := range in {
		left[i] = pair.left
		right[i] = pair.right
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range len(left) {
		result += int(math.Abs(float64(left[i] - right[i])))
	}

	return result
}
