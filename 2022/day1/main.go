package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Oh sh**. Don't Panic:", e)
	}
}

func part1(file *os.File) {
	var (
		numA  int
		numB  int
		count = 1
		elfId int
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			check(err)
			numB += num
		} else {
			if numB > numA {
				numA = numB
				elfId = count
			}
			numB = 0

			count++
		}
	}
	fmt.Printf("The elf carrying the most calroies is elf: %d\n", elfId)
	fmt.Printf("They are carrying %d calories\n", numA)
	err := scanner.Err()
	check(err)
}

func part2(file *os.File) {
	var (
		topElfBags []int
		totalCals  int
		newElf     int
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			check(err)
			newElf += num
		} else {
			topElfBags = append(topElfBags, newElf)
			sort.Ints(topElfBags)

			if len(topElfBags) > 3 {
				topElfBags = topElfBags[1:]
			}

			newElf = 0

		}
	}

	for _, v := range topElfBags {
		totalCals += v
	}
	fmt.Printf("Top three elves have %d calories\n", totalCals)
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
