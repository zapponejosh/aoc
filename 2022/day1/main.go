package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	elfBag []int
	numA   int
	numB   int
	count  = 1
	elfId  int
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			num, err := strconv.Atoi(line)
			check(err)
			elfBag = append(elfBag, num)
		} else {
			for _, v := range elfBag {
				numB += v
			}
			if numB > numA {
				numA = numB
				elfId = count
			}
			numB = 0

			count++
			elfBag = nil
		}
	}
	fmt.Printf("The elf carrying the most calroies is elf: %d\n", elfId)
	fmt.Printf("They are carrying %d calories\n", numA)
	err = scanner.Err()
	check(err)
}
