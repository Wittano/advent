package advent

import "testing"

func TestFirstDay(t *testing.T) {
	data := []NumberPair{
		{3, 4},
		{4, 3},
		{2, 5},
		{1, 3},
		{3, 9},
		{3, 3},
	}

	const expected = 11
	got := distanceBetweenSmallestNumbers(data)

	if got != expected {
		t.Fatalf("sum of distance between numbers is incorrect. exp %d, got %d", expected, got)
	}
}
