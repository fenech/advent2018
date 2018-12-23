package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checksum(input string) (bool, bool) {
	m := make(map[rune]int)

	for _, c := range input {
		m[c]++
	}

	var two, three bool
	for _, c := range m {
		if c == 2 {
			two = true
		}
		if c == 3 {
			three = true
		}
	}

	return two, three
}

func pt1(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)

	var twos, threes int
	for scanner.Scan() {
		two, three := checksum(scanner.Text())
		if two {
			twos++
		}
		if three {
			threes++
		}
	}

	return twos * threes
}

func removeDifferentCharacter(a string, b string) string {
	bs := make([]byte, 0, len(a)-1)

	for i := range a {
		if a[i] == b[i] {
			bs = append(bs, a[i])
		}
	}
	return string(bs)
}

func findCloseEnoughLineIndex(lines []string) int {
	target := lines[0]

	for i, l := range lines[1:] {
		idx := i + 1
		diff := 0
		for i := range l {
			if target[i] != l[i] {
				diff++
			}
			if diff > 1 {
				break
			}
		}

		if diff == 1 {
			return idx
		}
	}

	return -1
}

func pt2(handle io.Reader) string {
	scanner := bufio.NewScanner(handle)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for len(lines) > 2 {
		if i := findCloseEnoughLineIndex(lines); i > -1 {
			return removeDifferentCharacter(lines[0], lines[i])
		}
		lines = lines[1:]
	}

	return "no match"
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("part 1: %v\n", pt1(file))

	file.Seek(0, 0)

	fmt.Printf("part 2: %v\n", pt2(file))
}
