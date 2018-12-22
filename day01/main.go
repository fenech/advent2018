package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"io"
	"os"
	"strconv"
)

func pt1(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)

	c := 0
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			panic(err)
		}

		c += int(x)
	}
	return c
}

func pt2(handle io.Reader) int {
	scanner := bufio.NewScanner(handle)

	var r *ring.Ring
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			panic(err)
		}

		s := ring.New(1)
		s.Value = int(x)

		if r != nil {
			r = r.Prev().Link(s)
		} else {
			r = s
		}
	}

	seen := make(map[int]int)
	c := 0

	for {
		seen[c]++
		if seen[c] == 2 {
			return c
		}

		c += r.Value.(int)
		r = r.Next()
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("part 1: %v\n", pt1(file))

	file.Seek(0, 0)

	fmt.Printf("part 1: %v\n", pt2(file))
}
