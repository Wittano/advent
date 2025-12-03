package three

import (
	"fmt"
	"testing"
)

func TestBatteryPower_MaxPower(t *testing.T) {
	data := []struct {
		data   string
		result int
	}{
		{
			"987654321111111", 98,
		},
		{
			"811111111111119", 89,
		},
		{
			"234234234234278", 78,
		},
		{
			"818181911112111", 92,
		},
	}

	for _, in := range data {
		t.Run(fmt.Sprintf("Find max power in battery '%s'", in.data), func(t *testing.T) {
			battery, err := newBatteryPower(in.data)
			if err != nil {
				t.Fatal(err)
			}

			got, err := battery.MaxPower()
			if err != nil {
				t.Fatal(err)
			}

			if got != in.result {
				t.Fatalf("Invalid max got. want %d, got %d", in.result, got)
			}
		})
	}
}
