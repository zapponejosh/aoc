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

func findCommonItems(stringSacks ...string) string {
	sackCount := len(stringSacks)
	itemMap := make(map[string]int)
	// loop over each string set
	for i, strs := range stringSacks {
		var unique string
		for _, itm := range strs {
			if strings.Contains(unique, string(itm)) {
				continue
			}
			unique += string(itm)
			v, _ := itemMap[string(itm)]

			if v < i+1 {
				itemMap[string(itm)]++
			}

		}
	}
	// iterate over map
	for v := range itemMap {
		if itemMap[v] == sackCount {
			return v
		}
	}
	return ""
}

func part1(file *os.File) {
	var prioritySum int
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var commonItm string
		rucksack := scanner.Text()
		size := len(rucksack)
		compartment1 := rucksack[:(size / 2)]
		compartment2 := rucksack[(size / 2):]

		commonItm = findCommonItems(compartment1, compartment2)
		ItmPriority := strings.Index(priority, commonItm)
		prioritySum += (ItmPriority + 1)
	}
	fmt.Println("Part 1: ", prioritySum)
	err := scanner.Err()
	check(err)
}
func part2(file *os.File) {
	var prioritySum int
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	scanner := bufio.NewScanner(file)
	var count = 1
	var elves []string

	for scanner.Scan() {
		elf := scanner.Text()
		if count < 3 {
			elves = append(elves, elf)
			count++
		} else if count == 3 {
			elves = append(elves, elf)
			commonItem := findCommonItems(elves...)
			ItmPriority := strings.Index(priority, commonItem)
			prioritySum += (ItmPriority + 1)
			count = 1
			elves = nil
		}
	}
	fmt.Println("Part 2: ", prioritySum)
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
