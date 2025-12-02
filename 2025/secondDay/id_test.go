package secondDay

import (
	"fmt"
	"strings"
	"testing"
)

type testArg struct {
	data   string
	result int
}

func TestFindInvalidIds(t *testing.T) {
	data := []testArg{
		{
			"11-22", 2,
		},
		{
			"95-115", 1,
		},
		{
			"998-1012", 1,
		},
		{
			"1188511880-1188511890",
			1,
		},
		{
			"222220-222224",
			1,
		},
		{
			"1698522-1698528",
			0,
		},
		{
			"446443-446449",
			1,
		},
		{
			"38593856-38593862",
			1,
		},
		{"101-102", 0},
	}

	rawData := make([]string, len(data))
	expectedMap := make(map[string]int)
	for i, v := range data {
		rawData[i] = v.data
		expectedMap[v.data] = v.result
	}

	in := strings.Join(rawData, ",")

	res, err := FindInvalidIds(in)
	if err != nil {
		t.Fatal(err)
	}

	for key, value := range res {
		t.Run(fmt.Sprintf("find invalid id in %s range", key), func(t *testing.T) {
			if expectedMap[key] != value {
				t.Fatalf("Invalid number of invalid ids. want: %d, got: %d", expectedMap[key], value)
			}
		})
	}
}

func TestSumInvalidIds(t *testing.T) {
	data := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
	}

	sum, err := SumInvalidIDs(strings.Join(data, ","))
	if err != nil {
		t.Fatal(err)
	}

	if sum != 1227775554 {
		t.Fatalf("Invalid sum. want: %d, got: %d", 1227775554, sum)
	}
}

func TestSumTwoInvalidIds(t *testing.T) {
	data := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}

	sum, err := SumTwiceInvalidIDs(strings.Join(data, ","))
	if err != nil {
		t.Fatal(err)
	}

	const exp = 4174379265
	if sum != exp {
		t.Fatalf("Invalid sum. want: %d, got: %d", exp, sum)
	}
}
