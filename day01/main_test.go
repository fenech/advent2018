package main

import (
	"strings"
	"testing"
)

func Test_pt1(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{"+1\n+1\n+1\n"}, 3},
		{"second", args{"+1\n+1\n-2\n"}, 0},
		{"third", args{"-1\n-2\n-3\n"}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt1(strings.NewReader(tt.args.lines)); got != tt.want {
				t.Errorf("pt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pt2(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{"+1\n-1\n"}, 0},
		{"second", args{"+3\n+3\n+4\n-2\n-4\n"}, 10},
		{"third", args{"-6\n+3\n+8\n+5\n-6\n"}, 5},
		{"fourth", args{"+7\n+7\n-2\n-7\n-4\n"}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt2(strings.NewReader(tt.args.lines)); got != tt.want {
				t.Errorf("pt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
