package five

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	begin, end int
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.begin, r.end)
}

func extractRanges(scan *bufio.Scanner) ([]Range, error) {
	r := make([]Range, 0)

	for scan.Scan() {
		rawRange := scan.Text()
		if rawRange == "" {
			break
		}

		nums := strings.Split(rawRange, "-")
		if len(nums) != 2 {
			return nil, fmt.Errorf("five: invalid input range. Expected two numbers separated by '-': %q", nums)
		}

		begin, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}

		r = append(r, Range{begin, end})
	}

	return r, nil
}

func extractNumbers(scan *bufio.Scanner) ([]int, error) {
	nums := make([]int, 0)

	for scan.Scan() {
		num, err := strconv.Atoi(scan.Text())
		if err != nil {
			return nil, err
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func newRangePair(in string) ([]Range, []int, error) {
	scanner := bufio.NewScanner(strings.NewReader(in))

	ranges, err := extractRanges(scanner)
	if err != nil {
		return nil, nil, err
	}
	nums, err := extractNumbers(scanner)
	if err != nil {
		return nil, nil, err
	}

	return ranges, nums, nil
}

func CountFreshFood(in string) (amount uint, err error) {
	ranges, nums, err := newRangePair(in)
	if err != nil {
		return 0, err
	}

	for _, n := range nums {
		for _, r := range ranges {
			if n >= r.begin && n <= r.end {
				amount++
				break
			}
		}
	}

	return
}

func CountPossibleFreshFood(in string) (int, error) {
	ranges, _, err := newRangePair(in)
	if err != nil {
		return 0, err
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].begin < ranges[j].begin
	})

	var (
		amount  float64
		currMax float64
	)
	for _, r := range ranges {
		if float64(r.end) < currMax {
			continue
		}

		start := math.Max(float64(r.begin), currMax)
		currMax = math.Max(currMax, float64(r.end+1))
		amount += math.Abs(float64(r.end)-start) + 1
	}

	return int(amount), nil
}
