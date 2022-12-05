package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Oh sh**. Don't Panic:", e)
	}
}

func part1(file *os.File) {
	var (
		score int
	)
	value := map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissors
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		opp := string(line[0])
		self := string(line[2])
		score += value[self]

		switch {
		case value[opp] == value[self]: // Tie
			score += 3
		case self == "X" && opp == "C", self == "Y" && opp == "A", self == "Z" && opp == "B": // win
			score += 6
		default: // lose
		}

	}
	fmt.Println(score)
	err := scanner.Err()
	check(err)
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	part1(file)
}
