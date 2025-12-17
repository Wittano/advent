package main

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/five"
	"github.com/wittano/advent/text"
)

func TaskFivePartOne() {
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
