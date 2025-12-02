package firstDay

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RorationSequence string

func NewRorationSequence(in string) (result RorationSequence, err error) {
	if len(in) <= 1 {
		err = fmt.Errorf("day1: rotation sequence '%s' must include ", in)
		return
	}

	if in[0] != 'L' && in[0] != 'R' {
		err = fmt.Errorf("day1: rotation sequence '%s' must include 'L' or 'R' char", in)
		return
	}

	return RorationSequence(strings.TrimSpace(in)), nil
}

func (r RorationSequence) IsMoveLeft() bool {
	return r[0] == 'L'
}

func (r RorationSequence) Uint() (uint, error) {
	i, err := strconv.Atoi(string(r[1:]))
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

func UnlockPassword(start uint, moves []string) (uint, error) {
	if start > 99 {
		return 0, errors.New("day1: invalid start position. Number must be between 0 and 99")
	}

	var (
		password uint
		position = start
	)

	for _, move := range moves {
		sequence, err := NewRorationSequence(move)
		if err != nil {
			return 0, err
		}

		newMove, err := sequence.Uint()
		if err != nil {
			return 0, err
		}

		if sequence.IsMoveLeft() {
			position, _ = moveLeftPosition(position, newMove)
		} else {
			position, _ = moveRightPosition(position, newMove)
		}

		if position == 0 {
			password++
		}
	}

	return password, nil
}

func UnlockPasswordWithClick(start uint, moves []string) (uint, error) {
	if start > 99 {
		return 0, errors.New("day1: invalid start position. Number must be between 0 and 99")
	}

	var (
		password uint
		position = start
	)

	for _, move := range moves {
		sequence, err := NewRorationSequence(move)
		if err != nil {
			return 0, err
		}

		newMove, err := sequence.Uint()
		if err != nil {
			return 0, err
		}

		var click uint
		if sequence.IsMoveLeft() {
			position, click = moveLeftPosition(position, newMove)
		} else {
			position, click = moveRightPosition(position, newMove)
		}

		if position == 0 {
			password++
		} else {
			password += click
		}
	}

	return password, nil
}

const maxRotation = 100

func moveLeftPosition(current, move uint) (pos uint, click uint) {
	result := int(current) - int(move%maxRotation)

	if result < 0 {
		pos = uint(100 + result)
		if current != 0 {
			click++
		}
	} else {
		pos = uint(result)
	}

	if current != 0 {
		click += uint(math.Floor(float64(move / maxRotation)))
	}
	return
}

func moveRightPosition(current, move uint) (pos uint, click uint) {
	result := int(current + move%maxRotation)

	if result >= maxRotation {
		pos = uint(result % maxRotation)
		click++
	} else {
		pos = uint(result)
	}

	if current != 0 {
		click += uint(math.Floor(float64(move / maxRotation)))
	}
	return
}
