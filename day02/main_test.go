package main

import (
	"io"
	"strings"
	"testing"
)

func Test_checksum(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			"contains no letters that appear exactly two or three times",
			args{"abcdef"},
			false,
			false,
		},
		{
			"contains two a and three b, so it counts for both",
			args{"bababc"},
			true,
			true,
		},
		{
			"contains two b, but no letter appears exactly three times",
			args{"abbcde"},
			true,
			false,
		},
		{
			"contains three c, but no letter appears exactly two times",
			args{"abcccd"},
			false,
			true,
		},
		{
			"contains two a and two d, but it only counts once",
			args{"aabcdd"},
			true,
			false,
		},
		{
			"contains two e",
			args{"abcdee"},
			true,
			false,
		},
		{
			"contains three a and three b, but it only counts once",
			args{"ababab"},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checksum(tt.args.input)
			if got != tt.want {
				t.Errorf("checksum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checksum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
			"first",
			args{strings.NewReader("abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab\n")},
			12,
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

func Test_pt2(t *testing.T) {
	type args struct {
		handle io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example",
			args{strings.NewReader("abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz\n")},
			"fgij",
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

func Test_removeDifferentCharacter(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{"fghij", "fguij"}, "fgij"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDifferentCharacter(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("removeDifferentCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCloseEnoughLineIndex(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"return -1 when first line isn't close enough",
			args{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}},
			-1,
		},
		{
			"return index when first line is close enough",
			args{[]string{"fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCloseEnoughLineIndex(tt.args.lines); got != tt.want {
				t.Errorf("findCloseEnoughLineIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
