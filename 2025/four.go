package main

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/four"
	"github.com/wittano/advent/text"
)

func TaskFourPartOne() {
	rawContent, err := text.ReadFile("./2025/data/4.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := four.CountAvailablePaperRolls(rawContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 4, part 1: Result %d\n", sum)
}

func TaskFourPartTwo() {
	rawContent, err := text.ReadFile("./2025/data/4.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := four.CountAvailablePaperRollsWithRemove(rawContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 4, part 2: Result %d\n", sum)
}
