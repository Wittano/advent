package tasks

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/six"
	"github.com/wittano/advent/text"
)

func SixPartOne() {
	rawContent, err := text.ReadFile("./2025/data/6.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := six.SumHomeworks(rawContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 6, part 1: Result %d\n", int(sum))
}
