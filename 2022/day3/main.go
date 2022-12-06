package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Oh sh**. Don't Panic:", e)
	}
}

func part1(file *os.File) {
	var prioritySum int
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var commonItm string
		items := make(map[string]int)
		rucksack := scanner.Text()
		size := len(rucksack)
		// what if it's and odd number of items?
		compartment1 := rucksack[:(size / 2)]
		compartment2 := rucksack[(size / 2):]

		//could make this its own func
		for _, itm := range compartment1 {
			items[string(itm)] = 1 // should never be more than one because we only care if it's in both compartments
		}
		for _, v := range compartment2 {
			if _, ok := items[string(v)]; ok {
				commonItm = string(v)
			}
		}
		ItmPriority := strings.Index(priority, commonItm)
		prioritySum += (ItmPriority + 1)
	}
	fmt.Println(prioritySum)
	err := scanner.Err()
	check(err)
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	part1(file)
	// file.Seek(0, 0)
	// part2(file)
}
