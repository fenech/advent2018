package main

import (
	"testing"
)

func Test_pt1(t *testing.T) {
	type args struct {
		lines []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"first",
			args{[]int{1, 1, 1}},
			3,
		},
		{
			"second",
			args{[]int{1, 1, -2}},
			0,
		},
		{
			"third",
			args{[]int{-1, -2, -3}},
			-6,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt1(tt.args.lines); got != tt.want {
				t.Errorf("pt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pt2(t *testing.T) {
	type args struct {
		lines []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"first", args{[]int{+1, -1}}, 0},
		{"second", args{[]int{+3, +3, +4, -2, -4}}, 10},
		{"second", args{[]int{-6, +3, +8, +5, -6}}, 5},
		{"second", args{[]int{+7, +7, -2, -7, -4}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt2(tt.args.lines); got != tt.want {
				t.Errorf("pt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
