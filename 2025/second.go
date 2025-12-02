package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/wittano/advent/2025/secondDay"
)

func readWholeFile(filename string) (string, error) {
	f, err := os.Open("./2025/data/2.1.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	in, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(in), nil
}

func TaskTwoPartOne() {
	in, err := readWholeFile("./2025/data/2.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := secondDay.SumInvalidIDs(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2, part 1: %d\n", sum)
}

func TaskTwoPartTwo() {
	in, err := readWholeFile("./2025/data/2.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := secondDay.SumTwiceInvalidIDs(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 2, part 2: %d\n", sum)
}
