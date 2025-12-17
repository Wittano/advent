package five

import (
	"log"
	"testing"
)

func TestCountFreshFood(t *testing.T) {
	data := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	sum, err := CountFreshFood(data)
	if err != nil {
		log.Fatal(err)
	}

	const exp = 3
	if sum != exp {
		t.Errorf("invalid fresh food amount. got %d, want %d", sum, exp)
	}
}
