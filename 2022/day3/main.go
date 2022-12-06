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
	// sackCount := len(stringSacks)
	itemMap := make(map[string]int)
	// loop over each string set
	for i, strs := range stringSacks {
		var unique string
		for _, v := range strs {
			if !strings.Contains(unique, string(v)) {
				unique = unique + string(v)
			}
		}
		for _, itm := range unique {
			v, _ := itemMap[string(itm)]

			// if it's already in there we need to continue to the next string
			// but if it's there from a previous string we need to keep going!
			// what if i de-dup every string first?
			if v < i+1 {
				itemMap[string(itm)]++
			}

		}
	}
	// iterate over map
	// var commonItem string
	for v := range itemMap {
		if itemMap[v] == 3 {
			return v
		}
	}
	return "NOOOO"
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
			fmt.Println(commonItem)
			count = 1
			elves = nil
		}
	}
	fmt.Println(prioritySum)
	// fmt.Println(prioritySum)
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
