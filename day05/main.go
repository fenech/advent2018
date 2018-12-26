package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func react(polymer string) string {
	if polymer == "" {
		return polymer
	}

	for i := range polymer[:len(polymer)-1] {
		curr := rune(polymer[i])
		next := rune(polymer[i+1])
		if curr != next && unicode.ToUpper(curr) == unicode.ToUpper(next) {
			start := polymer[:i]
			var end string
			if i+2 < len(polymer) {
				end = polymer[i+2:]
			}
			return react(start + end)
		}
	}

	return polymer
}

func pt1(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)
	scanner.Scan()
	polymer := scanner.Text()
	return len(react(polymer))
}

func pt2(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)
	scanner.Scan()
	polymer := scanner.Text()

	product := react(polymer)
	units := make(map[rune]int)
	for i := range product {
		units[unicode.ToLower(rune(product[i]))] = 1
	}

	minLength := -1
	for c := range units {
		upperC := unicode.ToUpper(c)
		copy := make([]byte, 0, len(product))
		for i := range product {
			if rune(product[i]) != c && rune(product[i]) != upperC {
				copy = append(copy, product[i])
			}
		}

		reactedProduct := react(string(copy))
		if minLength == -1 || len(reactedProduct) < minLength {
			minLength = len(reactedProduct)
		}
	}

	return minLength
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
