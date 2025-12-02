package secondDay

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	regexp "github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

type Range struct {
	Start, End int
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.Start, r.End)
}

func NewRange(s string) (r Range, err error) {
	data := strings.Split(s, "-")

	if len(data) != 2 {
		return Range{}, errors.New("secondDay: invalid input")
	}

	r.Start, err = strconv.Atoi(data[0])
	if err != nil {
		return
	}

	r.End, err = strconv.Atoi(data[1])
	if err != nil {
		return
	}

	return
}

var regex = regexp.MustCompile(`^([0-9]+)\1$`, 0)

func FindInvalidIds(in string) (map[string]int, error) {
	invalidIdsCount := map[string]int{}
	for _, v := range strings.Split(in, ",") {
		r, err := NewRange(v)
		if err != nil {
			return nil, err
		}

		key := r.String()
		invalidIdsCount[key] = 0

		for i := r.Start; i <= r.End; i++ {
			value := strconv.Itoa(i)
			if len(value)%2 != 0 {
				continue
			}

			if regex.MatcherString(strconv.Itoa(i), 0).Matches() {
				invalidIdsCount[key] += i
			}
		}
	}

	return invalidIdsCount, nil
}

func SumInvalidIDs(in string) (sum int, err error) {
	for _, v := range strings.Split(in, ",") {
		r, err := NewRange(v)
		if err != nil {
			return 0, err
		}

		for i := r.Start; i <= r.End; i++ {
			value := strconv.Itoa(i)
			if len(value)%2 != 0 {
				continue
			}

			if regex.MatcherString(strconv.Itoa(i), 0).Matches() {
				sum += i
			}
		}
	}

	return
}

func SumTwiceInvalidIDs(in string) (sum int, err error) {
	var regex = regexp.MustCompile(`^(\d+)(\1)+$`, 0)
	for _, v := range strings.Split(in, ",") {
		r, err := NewRange(v)
		if err != nil {
			return 0, err
		}

		for i := r.Start; i <= r.End; i++ {
			if regex.MatcherString(strconv.Itoa(i), 0).Matches() {
				sum += i
			}
		}
	}

	return
}
