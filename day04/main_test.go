package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

const exampleInput = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

func Test_pt1(t *testing.T) {
	type args struct {
		handle io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example",
			args{strings.NewReader(exampleInput)},
			240,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt1(tt.args.handle); got != tt.want {
				t.Errorf("pt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want *record
	}{
		{
			"begin shift",
			args{"[1518-11-01 00:00] Guard #10 begins shift"},
			&record{time.Date(1518, time.November, 1, 0, 0, 0, 0, time.UTC), 10, BeginShift},
		},
		{
			"fall asleep",
			args{"[1518-11-01 00:05] falls asleep"},
			&record{time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC), Unknown, FallAsleep},
		},
		{
			"wake up",
			args{"[1518-11-01 00:25] wakes up"},
			&record{time.Date(1518, time.November, 1, 0, 25, 0, 0, time.UTC), Unknown, WakeUp},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pt2(t *testing.T) {
	type args struct {
		handle io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example",
			args{strings.NewReader(exampleInput)},
			4455,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt2(tt.args.handle); got != tt.want {
				t.Errorf("pt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
