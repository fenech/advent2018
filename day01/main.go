package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pt1(lines []int) int {
	c := 0
	for _, x := range lines {
		c += x
	}
	return c
}

func pt2(lines []int) int {
	seen := make(map[int]int)
	c := 0

	for {
		for _, x := range lines {
			seen[c]++
			if seen[c] == 2 {
				return c
			}
			c += x
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []int{}
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			panic(err)
		}
		lines = append(lines, int(x))
	}

	fmt.Printf("part 1: %v\n", pt1(lines))
	fmt.Printf("part 1: %v\n", pt2(lines))
}
