package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Oh sh**. Don't Panic:", e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	part1(file)
	file.Seek(0, 0)
	part2(file)

}

func part2(file *os.File) {
	var overlapCounter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		before, after, found := strings.Cut(row, ",")
		if !found {
			check(errors.New("separator not found"))
			panic("Dangit")
		}
		area1 := strings.Split(before, "-")
		area2 := strings.Split(after, "-")
		// area1[0] <= area2[0] && area2[0] <= area1[1]
		// area2[0] <= area1[0] && area1[0] <= area2[1]

		if (strToInt(area1[0]) <= strToInt(area2[0]) && strToInt(area2[0]) <= strToInt(area1[1])) || (strToInt(area2[0]) <= strToInt(area1[0]) && strToInt(area1[0]) <= strToInt(area2[1])) {
			overlapCounter += 1
			continue
		}

	}
	fmt.Println(overlapCounter)
}

func part1(file *os.File) {
	var containsCounter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		before, after, found := strings.Cut(row, ",")
		if !found {
			check(errors.New("separator not found"))
			panic("Dangit")
		}
		area1 := strings.Split(before, "-")
		area2 := strings.Split(after, "-")
		if strToInt(area1[0]) <= strToInt(area2[0]) && strToInt(area1[1]) >= strToInt(area2[1]) {
			containsCounter += 1
			continue // if the areas are identical
		}
		if strToInt(area2[0]) <= strToInt(area1[0]) && strToInt(area2[1]) >= strToInt(area1[1]) {
			containsCounter += 1
		}
	}
	fmt.Println(containsCounter)
}

func strToInt(s string) int {
	ret, err := strconv.Atoi(s)
	check(err)
	return ret
}
