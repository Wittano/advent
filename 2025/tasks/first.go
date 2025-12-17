package tasks

import (
	"fmt"
	"log"

	"github.com/wittano/advent/2025/firstDay"
	"github.com/wittano/advent/text"
)

func OnePartOne() {
	in, err := text.ReadFileByLines("./2025/data/1.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	result, err := firstDay.UnlockPassword(50, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First part result: %d\n", result)
}

func OnePartTwo() {
	in, err := text.ReadFileByLines("./2025/data/1.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	result, err := firstDay.UnlockPasswordWithClick(50, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Second part result: %d\n", result)
}
