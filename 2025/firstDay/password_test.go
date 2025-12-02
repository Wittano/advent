package firstDay

import (
	"fmt"
	"testing"
)

func TestFirstDayWithClick(t *testing.T) {
	type clickArgs struct {
		data   []string
		result uint
	}

	data := []clickArgs{
		{
			data: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			}, result: 6,
		},
		{
			data: []string{
				"R1000",
				"L30",
				"L15",
				"R150",
				"L426",
			},
			result: 15,
		},
		{
			data: []string{
				"L50",
				"L50",
			}, result: 1,
		},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("password '%d' for data %q", d.result, d.data), func(t *testing.T) {
			password, err := UnlockPasswordWithClick(50, d.data)
			if err != nil {
				t.Fatal(err)
			}

			if password != d.result {
				t.Errorf("password value is %d, want %d", password, d.result)
			}
		})
	}
}

func TestFirstDay(t *testing.T) {
	data := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	password, err := UnlockPassword(50, data)
	if err != nil {
		t.Fatal(err)
	}

	if password != 3 {
		t.Fatal("password is wrong. exp: 3, got:", password)
	}
}

type moveSequence struct {
	start, move, result uint
}

func TestMoveLeft(t *testing.T) {
	data := []moveSequence{
		{20, 10, 10},
		{0, 1, 99},
		{10, 20, 90},
		{10, 20, 90},
		{50, 100, 50},
		{50, 200, 50},
		{50, 210, 40},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("start on %d, move left %d and point on %d position", d.start, d.move, d.result), func(t *testing.T) {
			got, _ := moveLeftPosition(d.start, d.move)
			if got != d.result {
				t.Errorf("got %d, want %d", got, d.result)
			}
		})
	}
}

func TestMoveLeftCountClicks(t *testing.T) {
	data := []moveSequence{
		{20, 10, 0},
		{0, 1, 0},
		{1, 1, 0},
		{10, 20, 1},
		{50, 200, 2},
		{50, 210, 2},
		{50, 1000, 10},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("start on %d, move left %d and click %d times", d.start, d.move, d.result), func(t *testing.T) {
			_, got := moveLeftPosition(d.start, d.move)
			if got != d.result {
				t.Errorf("got %d, want %d", got, d.result)
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	data := []moveSequence{
		{20, 10, 30},
		{0, 1, 1},
		{99, 1, 0},
		{80, 30, 10},
		{80, 100, 80},
		{80, 110, 90},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("start on %d, move right %d and point on %d position", d.start, d.move, d.result), func(t *testing.T) {
			got, _ := moveRightPosition(d.start, d.move)
			if got != d.result {
				t.Errorf("got %d, want %d", got, d.result)
			}
		})
	}
}

func TestMoveRightWithClickCount(t *testing.T) {
	data := []moveSequence{
		{20, 10, 0},
		{0, 1, 0},
		{99, 1, 1},
		{80, 30, 1},
		{80, 100, 1},
		{80, 210, 2},
		{50, 1000, 10},
	}

	for _, d := range data {
		t.Run(fmt.Sprintf("start on %d, move right %d and click %d times", d.start, d.move, d.result), func(t *testing.T) {
			_, got := moveRightPosition(d.start, d.move)
			if got != d.result {
				t.Errorf("got %d, want %d", got, d.result)
			}
		})
	}
}
