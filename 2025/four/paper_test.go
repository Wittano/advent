package four

import "testing"

func TestCountAvailablePaperRolls(t *testing.T) {
	data := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	sum, err := CountAvailablePaperRolls(data)
	if err != nil {
		t.Fatal(err)
	}

	const exp = 13
	if sum != exp {
		t.Fatalf("invalid sum of available paper rolls: want %d, got %d", exp, sum)
	}
}
