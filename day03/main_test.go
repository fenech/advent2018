package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

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
			args{strings.NewReader("#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2\n")},
			4,
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
		want *claim
	}{
		{"first", args{"#1 @ 1,3: 4x4"}, &claim{1, start{x: 1, y: 3}, size{w: 4, h: 4}}},
		{"second", args{"#2 @ 3,1: 4x4"}, &claim{2, start{x: 3, y: 1}, size{w: 4, h: 4}}},
		{"third", args{"#3 @ 5,5: 2x2"}, &claim{3, start{x: 5, y: 5}, size{w: 2, h: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
