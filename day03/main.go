package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

type start struct {
	x int
	y int
}

type size struct {
	w int
	h int
}

type claim struct {
	id    int
	start start
	size  size
}

func parse(line string) *claim {
	re := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	res := re.FindStringSubmatch(line)
	p := func(s string) int {
		x, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			panic(err)
		}
		return int(x)
	}

	id, x, y, w, h := p(res[1]), p(res[2]), p(res[3]), p(res[4]), p(res[5])

	return &claim{
		id,
		start{x, y},
		size{w, h},
	}
}

func pt1(handle io.Reader) int {
	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		claim := parse(scanner.Text())
		for j := 0; j < claim.size.h; j++ {
			for i := 0; i < claim.size.w; i++ {
				y := claim.start.y + j
				x := claim.start.x + i

				if y >= len(fabric) {
					extra := make([][]int, y-len(fabric)+1)
					for i := range extra {
						extra[i] = make([]int, len(fabric[0]))
					}
					fabric = append(fabric, extra...)
				}

				if x >= len(fabric[0]) {
					for i := range fabric {
						extra := make([]int, x-len(fabric[0])+1)
						fabric[i] = append(fabric[i], extra...)
					}
				}

				fabric[y][x]++
			}
		}
	}

	overlap := 0

	for j := range fabric {
		for i := range fabric[j] {
			if fabric[j][i] > 1 {
				overlap++
			}
		}
	}

	return overlap
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("part 1: %v\n", pt1(file))

	file.Seek(0, 0)
}
