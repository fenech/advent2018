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
	return 0
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
