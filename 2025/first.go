package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/wittano/advent/2025/firstDay"
)

func readInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	raw := bufio.NewScanner(f)
	in := make([]string, 0, len(raw.Bytes()))
	for raw.Scan() {
		in = append(in, raw.Text())
	}

	return in, nil
}

func firstPart() {
	in, err := readInput("./2025/data/1.1.txt")
	if err != nil {
		log.Fatal(err)
	}

	result, err := firstDay.UnlockPassword(50, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("First part result: %d\n", result)
}

func secondPart() {
	in, err := readInput("./2025/data/1.2.txt")
	if err != nil {
		log.Fatal(err)
	}

	result, err := firstDay.UnlockPasswordWithClick(50, in)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Second part result: %d\n", result)
}

func main() {
	firstPart()
	secondPart()
}
