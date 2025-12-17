package tasks

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/five"
	"github.com/wittano/advent/text"
)

func FivePartOne() {
	rawContent, err := text.ReadFile("./2025/data/5.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := five.CountFreshFood(rawContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 5, part 1: Result %d\n", sum)
}

func FivePartTwo() {
	rawContent, err := text.ReadFile("./2025/data/5.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := five.CountPossibleFreshFood(rawContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 5, part 2: Result %d\n", sum)
}
