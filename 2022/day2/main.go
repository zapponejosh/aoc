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
func part2(file *os.File) {
	var (
		score int
	)
	value := map[string]int{
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors
		"X": 0, // lost
		"Y": 3, // draw
		"Z": 6, // win
	}

	// X is Lose
	// Y is draw
	// Z is win

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		opp := string(line[0])
		result := string(line[2])
		score += value[result]

		if result == "Y" { // draw
			score += value[opp]
		} else if result == "Z" { // win
			switch opp {
			case "A":
				score += 2 // Paper beats rock
			case "B":
				score += 3 // Scissors beats paper
			case "C":
				score += 1 // rock beats scissors
			}
		} else { // lose
			switch opp {
			case "A":
				score += 3 // Rock beats scissors
			case "B":
				score += 1 // Paper beats rock
			case "C":
				score += 2 // Scissors beats paper
			}
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
	file.Seek(0, 0)
	part2(file)
}
