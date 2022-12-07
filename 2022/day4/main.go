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
		area1, area2, err := parseRow(row)
		check(err)

		if (area1[0] <= area2[0] && area2[0] <= area1[1]) || (area2[0] <= area1[0] && area1[0] <= area2[1]) {
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
		area1, area2, err := parseRow(row)
		check(err)
		if area1[0] <= area2[0] && area1[1] >= area2[1] {
			containsCounter += 1
			continue // if the areas are identical
		}
		if area2[0] <= area1[0] && area2[1] >= area1[1] {
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

func parseRow(row string) ([]int, []int, error) {
	before, after, found := strings.Cut(row, ",")
	if !found {
		return nil, nil, errors.New("Separator not found")
	}
	firstS := strings.Split(before, "-")
	secondS := strings.Split(after, "-")

	firstI, err := convertToInts(firstS)
	if err != nil {
		return nil, nil, err
	}
	secondI, err := convertToInts(secondS)
	if err != nil {
		return nil, nil, err
	}
	return firstI, secondI, nil
}
func convertToInts(s []string) ([]int, error) {
	var ret []int
	for _, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ret = append(ret, num)

	}
	return ret, nil
}
