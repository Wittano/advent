package four

import (
	"bufio"
	"fmt"
	"math/bits"
	"regexp"
	"strings"
)

type PaperLineValue uint8

const (
	None PaperLineValue = 1 << iota
	Dot
	Paper
)

type PaperRolls [][]PaperLineValue

func (p PaperRolls) String() string {
	var b strings.Builder

	for _, line := range p {
		for _, v := range line {
			switch v {
			case Dot:
				b.WriteString(".")
			case Paper:
				b.WriteString("@")
			case None:
				b.WriteString(".")
			}
		}
	}

	return b.String()
}

func (p PaperRolls) LineValue(x, y int) (val PaperLineValue) {
	defer func() {
		if err := recover(); err != nil {
			val = None
		}
	}()

	val = p[y][x]

	return
}

func (p PaperRolls) IsPaperAccess(x, y int) bool {
	if p.LineValue(x, y) != Paper {
		return false
	}

	const (
		secondPosition = 1
		thirdPosition  = 2
	)

	top := p.LineValue(x-1, y-1)&Paper | (p.LineValue(x, y-1)&Paper)<<secondPosition | (p.LineValue(x+1, y-1)&Paper)<<thirdPosition
	middle := p.LineValue(x-1, y)&Paper | (p.LineValue(x+1, y)&Paper)<<secondPosition
	bottom := p.LineValue(x-1, y+1)&Paper | (p.LineValue(x, y+1)&Paper)<<secondPosition | (p.LineValue(x+1, y+1)&Paper)<<thirdPosition

	return bits.OnesCount(uint(top))+bits.OnesCount8(uint8(middle))+bits.OnesCount(uint(bottom)) < 4
}

var lineReg = regexp.MustCompile(`^[.@]+$`)

func NewPaperRolls(rawSheet string) (rolls PaperRolls, err error) {
	rolls = make(PaperRolls, 0, strings.Count(rawSheet, "\n"))

	scan := bufio.NewScanner(strings.NewReader(rawSheet))
	for scan.Scan() {
		rawLine := scan.Text()

		if !lineReg.MatchString(rawLine) {
			return nil, fmt.Errorf("four: invalid rawLine input: %s", rawLine)
		}

		line := make([]PaperLineValue, len(rawLine))

		for i, c := range rawLine {
			switch c {
			case '.':
				line[i] = Dot
			case '@':
				line[i] = Paper
			default:
				line[i] = None
			}
		}

		rolls = append(rolls, line)
	}

	return
}

func CountAvailablePaperRolls(rawLines string) (sum uint, err error) {
	rolls, err := NewPaperRolls(rawLines)
	if err != nil {
		return 0, err
	}

	for y, line := range rolls {
		for x, mark := range line {
			if mark == Paper && rolls.IsPaperAccess(x, y) {
				sum++
			}
		}
	}

	return
}

func CountAvailablePaperRollsWithRemove(rawLines string) (sum uint, err error) {
	rolls, err := NewPaperRolls(rawLines)
	if err != nil {
		return 0, err
	}

	type position struct {
		x, y int
	}

	eliminatedPapers := make([]position, 0, len(rolls))
	for wasRemoved := true; wasRemoved; {
		for y, line := range rolls {
			for x, mark := range line {
				if mark == Paper && rolls.IsPaperAccess(x, y) {
					sum++
					eliminatedPapers = append(eliminatedPapers, position{x, y})
				}
			}
		}

		if len(eliminatedPapers) == 0 {
			wasRemoved = false
		}

		for _, pos := range eliminatedPapers {
			rolls[pos.y][pos.x] = None
		}

		eliminatedPapers = eliminatedPapers[:0]
	}

	return
}
