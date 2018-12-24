package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Event is an enum representing the event that a record corresponds to
type Event int

const (
	BeginShift Event = 0
	FallAsleep Event = 1
	WakeUp     Event = 2
)

type Guard int

const (
	Unknown Guard = -1
)

type record struct {
	time  time.Time
	guard Guard
	event Event
}

func parseTime(value string) time.Time {
	time, err := time.Parse("[2006-01-02 15:04]", value)
	if err != nil {
		panic(err)
	}

	return time
}

func parseEvent(value string) (Guard, Event) {
	if value == "falls asleep" {
		return Unknown, FallAsleep
	}
	if value == "wakes up" {
		return Unknown, WakeUp
	}

	re := regexp.MustCompile("Guard #(\\d+) begins shift")
	m := re.FindStringSubmatch(value)

	g, err := strconv.ParseInt(m[1], 10, 0)
	if err != nil {
		panic(err)
	}

	return Guard(g), BeginShift
}

func parse(line string) *record {
	time := parseTime(line[:18])
	id, event := parseEvent(line[19:])
	return &record{time, id, event}
}

func pt1(handle io.Reader) int {
	var id int
	return id
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
