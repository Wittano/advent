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

func TestCountPossibleFreshFood(t *testing.T) {
	data := []struct {
		got string
		exp int
	}{
		{
			`3-5
10-14
16-20
12-18
`, 14,
		},
		{
			`1-5
10-15
17-18
7-10
9-10
5-6
40-50`, 28,
		},
	}

	for _, d := range data {
		t.Run(d.got, func(t *testing.T) {
			sum, err := CountPossibleFreshFood(d.got)
			if err != nil {
				log.Fatal(err)
			}

			if sum != d.exp {
				t.Errorf("invalid fresh food amount. got %d, want %d", sum, d.exp)
			}
		})
	}
}
