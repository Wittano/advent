package tasks

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/secondDay"
	"github.com/wittano/advent/text"
)

func TwoPartOne() {
	in, err := text.ReadFile("./2025/data/2.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := secondDay.SumInvalidIDs(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2, part 1: %d\n", sum)
}

func TwoPartTwo() {
	in, err := text.ReadFile("./2025/data/2.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := secondDay.SumTwiceInvalidIDs(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2, part 2: %d\n", sum)
}
