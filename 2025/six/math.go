package six

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type MathOperation byte

const (
	Plus MathOperation = iota
	Minus
	Multiply
	Divide
)

func newMathOperation(in string) (op MathOperation) {
	switch in {
	case "+":
		op = Plus
	case "-":
		op = Minus
	case "*":
		op = Multiply
	case "/":
		op = Divide
	default:
		log.Fatalf("six: unknown math operation: '%s'", in)
	}

	return
}

type Homework struct {
	Op   MathOperation
	Nums []int
}

func (h Homework) Operation() (sum float64) {
	if h.Op == Multiply {
		sum = 1
	}

	for _, num := range h.Nums {
		switch h.Op {
		case Plus:
			sum += float64(num)
		case Multiply:
			sum *= float64(num)
		default:
			log.Fatalf("six: unknown operation: '%q'", h.Op)
		}
	}

	return
}

func newHomework(in string) ([]Homework, error) {
	r := strings.NewReader(strings.TrimSpace(in))
	s := bufio.NewScanner(r)

	firstLine := extractStrings(s.Text())
	homeworks := make([]Homework, 0, len(firstLine))
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	numsCount := strings.Count(in, "\n")
	for rowIdx := 0; s.Scan(); rowIdx++ {
		words := extractStrings(s.Text())
		if len(homeworks) != len(words) {
			for i := len(homeworks); i < len(words); i++ {
				homeworks = append(homeworks, Homework{
					Nums: make([]int, numsCount),
				})
			}
		}

		for i, line := range words {
			if checkOperationChar(line) {
				homeworks[i].Op = newMathOperation(line)
			} else {
				num, err := strconv.Atoi(line)
				if err != nil {
					return nil, err
				}

				homeworks[i].Nums[rowIdx] = num
			}
		}
	}

	return homeworks, nil
}

func extractStrings(line string) []string {
	trimLine := strings.TrimSpace(line)
	var builder strings.Builder

	words := make([]string, 0, len(trimLine))
	for _, c := range trimLine {
		if c != ' ' {
			builder.WriteRune(c)
		} else if builder.Len() > 0 {
			words = append(words, builder.String())
			builder.Reset()
		}
	}

	if builder.Len() > 0 {
		words = append(words, builder.String())
	}

	return words
}

func checkOperationChar(opChar string) bool {
	return opChar == "+" || opChar == "-" || opChar == "*" || opChar == "/"
}

func SumHomeworks(in string) (sum float64, err error) {
	homeworks, err := newHomework(in)
	if err != nil {
		return 0, err
	}

	for _, homework := range homeworks {
		sum += homework.Operation()
	}

	return sum, nil
}
