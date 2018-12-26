package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
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

type SleepState int

type Guard struct {
	ID         GuardID
	Minutes    [60]int
	MaxMinute  int
	TotalSleep int
}

type record struct {
	time  time.Time
	ID    GuardID
	event Event
}

func parseTime(value string) time.Time {
	time, err := time.Parse("[2006-01-02 15:04]", value)
	if err != nil {
		panic(err)
	}

	return time
}

type GuardID int

const Unknown GuardID = -1

func parseEvent(value string) (GuardID, Event) {
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

	return GuardID(g), BeginShift
}

func parse(line string) *record {

	time := parseTime(line[:18])
	guard, event := parseEvent(line[19:])

	return &record{time, guard, event}
}

func createRecords(handle io.Reader) []*record {
	records := []*record{}

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		records = append(records, parse(scanner.Text()))
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].time.Before(records[j].time)
	})

	return records
}

func pt1(handle io.Reader) int {
	records := createRecords(handle)

	var guard *Guard
	var maxGuard *Guard

	guards := make(map[GuardID]*Guard)
	for i, r := range records {
		if r.event == BeginShift {
			if r.ID == Unknown {
				panic("unknown guard")
			}

			if guards[r.ID] == nil {
				guards[r.ID] = &Guard{
					ID: r.ID,
				}
			}

			guard = guards[r.ID]
		} else if r.event == FallAsleep {
			if records[i+1].event != WakeUp {
				panic("expected wake up after fall asleep")
			}
			start, end := r.time, records[i+1].time
			for t := start; t.Before(end); t = t.Add(time.Minute) {
				minute := t.Minute()
				guard.Minutes[minute]++
				if guard.Minutes[minute] > guard.Minutes[guard.MaxMinute] {
					guard.MaxMinute = minute
				}
				guard.TotalSleep++
				if maxGuard == nil || guard.TotalSleep > maxGuard.TotalSleep {
					maxGuard = guard
				}
			}
		}
	}

	return int(maxGuard.ID) * maxGuard.MaxMinute
}

func pt2(handle io.Reader) int {
	records := createRecords(handle)

	var guard *Guard
	var maxGuard *Guard

	guards := make(map[GuardID]*Guard)
	for i, r := range records {
		if r.event == BeginShift {
			if r.ID == Unknown {
				panic("unknown guard")
			}

			if guards[r.ID] == nil {
				guards[r.ID] = &Guard{
					ID: r.ID,
				}
			}

			guard = guards[r.ID]
		} else if r.event == FallAsleep {
			if records[i+1].event != WakeUp {
				panic("expected wake up after fall asleep")
			}
			start, end := r.time, records[i+1].time
			for t := start; t.Before(end); t = t.Add(time.Minute) {
				minute := t.Minute()
				guard.Minutes[minute]++
				if guard.Minutes[minute] > guard.Minutes[guard.MaxMinute] {
					guard.MaxMinute = minute
				}
				guard.TotalSleep++
				if maxGuard == nil || guard.Minutes[guard.MaxMinute] > maxGuard.Minutes[maxGuard.MaxMinute] {
					maxGuard = guard
				}
			}
		}
	}

	return int(maxGuard.ID) * maxGuard.MaxMinute
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
