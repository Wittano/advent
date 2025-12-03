package main

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/three"
	"github.com/wittano/advent/text"
)

func TaskThreePartOne() {
	in, err := text.ReadFileByLines("./2025/data/3.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum, err := three.SumBatteryPower(in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Day 3 part 1, Result: %d\n", sum)
}
